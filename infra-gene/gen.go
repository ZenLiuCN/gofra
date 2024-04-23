package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/ZenLiuCN/goinfra/modeler"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "show version",
		Aliases: []string{"v"},
	}
	err := (&cli.App{
		UseShortOptionHandling: true,
		Name:                   "Entity Generator",
		Version:                "v0.0.1",
		Usage:                  "Generate entity components for go-infra",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "enum", Aliases: []string{"e"}, Usage: "The begin field enum number,must greater than 8", DefaultText: "9"},
			&cli.StringSliceFlag{
				Name:     "type",
				Usage:    "type names of current file or directory",
				Required: true,
				Aliases:  []string{"t"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "debug generator",
				Aliases: []string{"d"},
			},
			&cli.StringSliceFlag{
				Name:     "tags",
				Usage:    "tags to apply",
				Required: false,
				Aliases:  []string{"g"},
			},
			&cli.StringFlag{
				Name:        "out",
				DefaultText: "infra_entities.go",
				Usage:       "output file name",
				Required:    false,
				Value:       "",
				Aliases:     []string{"o"},
			},
			&cli.BoolFlag{
				Name:        "accessor",
				Usage:       "generate accessor",
				Aliases:     []string{"a"},
				Value:       false,
				DefaultText: "false",
			},
		},

		Suggest:              true,
		EnableBashCompletion: true,
		Action: func(c *cli.Context) error {
			var file []string
			if c.Args().Len() == 0 {
				file = append(file, ".")
			} else {
				file = append(file, c.Args().Slice()...)
			}
			tags := c.StringSlice("tags")
			var dir string
			if len(file) == 1 && isDir(file[0]) {
				dir = file[0]
			} else if len(tags) != 0 {
				log.Fatal("--tags can only applies with directory")
			} else {
				dir = filepath.Dir(file[0])
			}
			g := new(Generator)
			{
				g.dir = dir
				g.tags = tags
				g.files = file
				g.types = c.StringSlice("type")
				g.access = c.Bool("accessor")
				g.index = modeler.FIELD(c.Int("enum"))
				g.out = c.String("out")
				if c.Bool("debug") {
					g.Logf = log.Printf
				}
			}
			g.generate()
			return nil
		},
	}).Run(os.Args)
	if err != nil {
		panic(err)
	}
}
func isDir(name string) bool {
	if info, err := os.Stat(name); err != nil {
		log.Fatal(err)
	} else {
		return info.IsDir()
	}
	return false
}

type Generator struct {
	dir    string
	tags   []string
	files  []string
	types  []string
	access bool
	index  modeler.FIELD
	Context
	declEntities    []*Entity
	out             string
	seq             modeler.FIELD
	enumWriter      *Writer
	initWriter      *Writer
	variableWriter  *Writer
	factoriesWriter *Writer
}

func (g *Generator) generate() {
	g.Parse(g.tags, g.files)
	g.Process()
}

func (g *Generator) Process() {
	for _, p := range g.Pkg {
		for _, file := range p.Files {
			if file.file != nil {
				ast.Inspect(file.file, decl(file, p, g))
			}
		}
	}
	for _, declEntity := range g.declEntities {
		declEntity.process()
	}
	g.Write()
}
func (g *Generator) Write() {
	if g.out == "" {
		g.out = "infra_entities.go"
	}
	if g.index == 0 {
		g.index = modeler.FIELD_BUILTIN_MAX + 1
	}
	g.seq = g.index
	g.enumWriter = NewWriter().F(`
const (
`)
	g.initWriter = NewWriter().F("\nfunc init(){\n")
	g.variableWriter = NewWriter().F("\nvar (\n")
	g.factoriesWriter = NewWriter()
	ent := NewWriter().Import("context")
	for _, entity := range g.declEntities {
		entity.write(g)
		ent.F(`
type %[1]sEntity struct{
	%[1]s
	modeler.BaseEntity[%[2]s,%[1]s]
}
func new%[1]sEntity(tab string,configurer modeler.Configurer, s %[1]s,executor modeler.Executor) (e *%[1]sEntity){
	e=&%[1]sEntity{
		%[1]s:s,
	}
	e.BaseEntity = modeler.NewBaseEntity[%[2]s,%[1]s](
			configurer,
			&e.%[1]s,
			%[1]sFields,
			executor,
			modeler.NewBaseSQLMaker[%[2]s, %[1]s](tab, configurer, %[1]sFields),
			func(ctx context.Context) bool {
				return true
			},
		)
	return
}
`, entity.Name, entity.IdType)
	}
	g.enumWriter.F(")\n").
		Append(g.variableWriter.F(")\n")).
		Append(g.initWriter.F("}\n")).
		Append(ent)
	out := NewWriter().F("package %s\n", g.Pkg[0].Name).F(`
import (
`)
	for s := range g.enumWriter.Imports {
		out.F("\"%s\"\n", s)
	}
	out.F("\t)\n")
	out.Append(g.enumWriter)
	src, err := format.Source(out.buf.Bytes())
	if err != nil {
		log.Printf("warning: invalid generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		fn.Panic(os.WriteFile(g.out, out.buf.Bytes(), os.ModePerm))
	} else {
		fn.Panic(os.WriteFile(g.out, src, os.ModePerm))
	}

}

type Entity struct {
	gen  *Generator
	pkg  *Package
	file *File

	typeSpec   *ast.TypeSpec
	structType *ast.StructType

	Name   string
	Fields []*Field
	IdType string
}

func (e *Entity) foundSelectorType(x *ast.SelectorExpr) types.Object {
	return e.pkg.Pkg.Types.Imports()[slices.IndexFunc(e.pkg.Pkg.Types.Imports(), func(p *types.Package) bool {
		return p.Name() == x.X.(*ast.Ident).Name
	})].Scope().Lookup(x.Sel.Name)
}
func (e *Entity) findIndexType(x *ast.IndexExpr) types.Object {
	return e.pkg.Pkg.TypesInfo.Defs[x.X.(*ast.Ident)]
}
func (e *Entity) process() {
	for _, field := range e.structType.Fields.List {
		switch x := field.Type.(type) {
		case *ast.Ident:
			def := e.pkg.Pkg.TypesInfo.Defs[x]
			e.pkg.Printf("%s %#+v\n", x, def)
			if def == nil {
				for _, name := range field.Names {
					e.Fields = append(e.Fields, (&Field{
						Entity:    e,
						Name:      name.Name,
						Tags:      NewTagsOf(field.Tag),
						TypeName:  x.Name,
						IdentType: x,
					}).parseColumn())
				}
			} else {
				panic(def)
			}
		case *ast.IndexExpr:
			def := e.findIndexType(x)
			e.pkg.Printf("%#+v\n", def)
			if len(field.Names) == 0 {
				if v, ok := def.(*types.Var); ok && v.Embedded() && v.IsField() {
					if tp, ok := v.Type().Underlying().(*types.Struct); ok {
						e.pkg.Printf("%#+v", tp)
						f := recurringFields(e, tp, x, nil)
						e.Fields = append(e.Fields, f...)
						continue
					}
				}
				panic(def)
			} else {
				for _, name := range field.Names {
					e.Fields = append(e.Fields, (&Field{
						Entity:   e,
						Name:     name.Name,
						Column:   "",
						Index:    0,
						Tags:     NewTagsOf(field.Tag),
						VarField: nil,
						TypeName: def.Name(),
					}).parseColumn())
				}
				panic(e.Fields)
			}
		case *ast.SelectorExpr:
			p := e.foundSelectorType(x)
			e.pkg.Printf("%#+v\n", p)
			if v, ok := p.(*types.TypeName); ok {
				if len(field.Names) == 0 {
					if tp, ok := v.Type().Underlying().(*types.Struct); ok {
						e.pkg.Printf("%#+v", tp)
						f := recurringFields(e, tp, nil, x)
						e.Fields = append(e.Fields, f...)
						continue
					}
				} else {
					if tp, ok := v.Type().Underlying().(*types.Struct); ok {
						e.pkg.Printf("%#+v", tp)
						for _, name := range field.Names {
							e.Fields = append(e.Fields, (&Field{
								Entity:       e,
								Name:         name.Name,
								Column:       "",
								Index:        0,
								Tags:         NewTagsOf(field.Tag),
								VarField:     nil,
								StructType:   tp,
								TypeName:     x.X.(*ast.Ident).Name + "." + x.Sel.Name,
								SelectorType: x,
							}).parseColumn())
						}
						continue
					}
				}
			}
			panic(p)
		case nil:
			continue
		default:
			panic(x)
		}
	}
}

func (e *Entity) write(g *Generator) {

	getter := NewWriter().Import("github.com/ZenLiuCN/goinfra/modeler").Import("fmt").Import("time").Import("database/sql")
	setter := NewWriter()
	naming := NewWriter()
	getter.F(`
func(u *%s,f modeler.FIELD)any{
	switch(f){
	case modeler.FIELD_ID:
		return u.Id
	case modeler.FIELD_CREATE_AT:
		return u.CreateAt
	case modeler.FIELD_MODIFIED_AT:
		return u.ModifiedAt
	case modeler.FIELD_REMOVED:
		return u.Removed
	case modeler.FIELD_VERSION:
		return u.Version
	case modeler.FIELD_CREATE_BY:
		return u.CreateBy
	case modeler.FIELD_MODIFIED_BY:
		return u.ModifiedBy
`, e.Name)
	setter.F(`func(u *%s,f modeler.FIELD,v any){
switch(f){
case modeler.FIELD_ID:
	if x, ok := v.(int64); ok {
		u.Id = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_CREATE_AT:
	if x, ok := v.(time.Time); ok {
		u.CreateAt = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_MODIFIED_AT:
	if x, ok := v.(time.Time); ok {
		u.ModifiedAt = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_REMOVED:
	if x, ok := v.(bool); ok {
		u.Removed = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_VERSION:
	if x, ok := v.(int); ok {
		u.Version = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_CREATE_BY:
	if x, ok := v.(int64); ok {
		u.CreateBy = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
case modeler.FIELD_MODIFIED_BY:
	if x, ok := v.(int64); ok {
		u.ModifiedBy = x
	} else {
		panic(fmt.Errorf("bad field %%T type of %%d", v, f))
	}
`, e.Name)
	naming.F(`map[modeler.FIELD]string{
			modeler.FIELD_ID:          "id",
			modeler.FIELD_CREATE_AT:   "create_at",
			modeler.FIELD_MODIFIED_AT: "modified_at",
			modeler.FIELD_REMOVED:     "removed",
			modeler.FIELD_VERSION:     "version",
			modeler.FIELD_CREATE_BY:   "create_by",
			modeler.FIELD_MODIFIED_BY: "modified_by",
`)
	for _, field := range e.Fields {
		switch {
		case field.Column == "id" || field.Name == "Id":
			field.Index = modeler.FIELD_ID
			e.IdType = field.TypeName
		case field.Column == "create_at" || field.Name == "CreateAt":
			field.Index = modeler.FIELD_CREATE_AT
		case field.Column == "removed" || field.Name == "Removed":
			field.Index = modeler.FIELD_REMOVED
		case field.Column == "version" || field.Name == "Version":
			field.Index = modeler.FIELD_VERSION
		case field.Column == "modified_at" || field.Name == "ModifiedAt":
			field.Index = modeler.FIELD_MODIFIED_AT
		case field.Column == "create_by" || field.Name == "CreateBy":
			field.Index = modeler.FIELD_CREATE_BY
		case field.Column == "modified_by" || field.Name == "ModifiedBy":
			field.Index = modeler.FIELD_MODIFIED_BY
		default:
			field.Index = g.seq
			g.enumWriter.F(`
//%[1]s%[2]s generate for %[1]s.%[2]s <%[4]s::%[5]s>
%[1]s%[2]s modeler.FIELD = %[3]d
`, e.Name, field.Name, field.Index, field.Column, field.TypeName)
			naming.F("%s%s:\"%s\",\n", e.Name, field.Name, field.Column)
			getter.F(`
case %[1]s%[2]s:
		return u.%[2]s
`, e.Name, field.Name)
			setter.F(`
case %[1]s%[2]s:
`, e.Name, field.Name)
			switch field.TypeName {
			default:
				setter.F(`
if x,ok:=v.(%s);ok{
	u.%s=x
} else {
	panic(fmt.Errorf("bad field %%T type of %%d", v, f))
}`, field.TypeName, field.Name)
			}
			g.seq += 1
		}
	}
	if e.IdType == "" {
		panic(fmt.Errorf("not found IdType for %s", e.Name))
	}
	g.variableWriter.F("\n%[1]sFields modeler.EntityInfo[%[2]s, %[1]s]\n", e.Name, e.IdType)
	g.initWriter.F("\n%[1]sFields=modeler.EntityInfoBuilder[%[2]s,%[1]s](\n", e.Name, e.IdType)
	g.initWriter.Append(getter.F(`
default:
	panic(fmt.Errorf("invalid field %%d", f))
}},
`))
	g.initWriter.Append(setter.F(`
default:
	panic(fmt.Errorf("invalid field %%d", f))
}},
`))
	g.initWriter.Append(naming.F(`},
`))
	g.initWriter.F("\n)")
}
func recurringFields(e *Entity, t *types.Struct, index *ast.IndexExpr, sel *ast.SelectorExpr) (r []*Field) {
	n := t.NumFields()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		if f.Embedded() {
			r = append(r, recurringFields(e, f.Type().Underlying().(*types.Struct), index, sel)...)
		} else {
			r = append(r, (&Field{
				Entity:       e,
				Name:         f.Name(),
				Column:       "",
				Index:        0,
				VarField:     f,
				Tags:         NewTags(t.Tag(i)),
				TypeName:     f.Type().String(),
				IndexType:    index,
				SelectorType: sel,
			}).parseColumn())
		}
	}
	return r
}

type Field struct {
	Entity       *Entity
	Name         string
	Column       string
	Index        modeler.FIELD
	Tags         Tags
	VarField     *types.Var
	StructType   *types.Struct
	TypeName     string
	IdentType    *ast.Ident
	IndexType    *ast.IndexExpr
	SelectorType *ast.SelectorExpr
}

func (f *Field) parseColumn() *Field {
	if f.Tags == nil {
		f.Column = f.Name
	} else if v, ok := f.Tags["db"]; ok && len(v) > 0 {
		f.Column = v[0]
	} else {
		f.Column = f.Name
	}
	return f
}

func decl(f *File, p *Package, g *Generator) func(node ast.Node) bool {
	return func(node ast.Node) bool {
		dec, ok := node.(*ast.GenDecl)
		if !ok || dec.Tok != token.TYPE {
			return true
		}
		for _, spec := range dec.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				if st, ok := ts.Type.(*ast.StructType); ok {
					if slices.Contains(g.types, ts.Name.Name) {
						g.declEntities = append(g.declEntities, &Entity{
							Name:       ts.Name.Name,
							gen:        g,
							pkg:        p,
							file:       f,
							typeSpec:   ts,
							structType: st,
						})
					}

				}
			}
		}
		return true
	}
}
