package main

//https://cs.opensource.google/go/x/tools/+/master:cmd/stringer/stringer.go;drc=daf94608b5e2caf763ba634b84e7a5ba7970e155;l=382
import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"github.com/urfave/cli/v2"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "show cli version",
		Aliases: []string{"v"},
	}
	err := (&cli.App{
		UseShortOptionHandling: true,
		Name:                   "Entity Generator",
		Version:                "v0.0.1",
		Usage:                  "Generate entity components for go-infra",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "enum", Aliases: []string{"e"}, Value: 9, Usage: "The begin field enum number,must greater than 8", DefaultText: "9"},
			&cli.StringSliceFlag{
				Name:     "type",
				Usage:    "type names of current file or directory",
				Required: true,
				Aliases:  []string{"t"},
			},
			&cli.StringSliceFlag{
				Name:     "tags",
				Usage:    "tags to apply",
				Required: false,
				Aliases:  []string{"g"},
			},
			&cli.StringFlag{
				Name:        "out",
				DefaultText: "<type>_entity.go",
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
			g := &Generator{
				dir:    dir,
				tags:   tags,
				files:  file,
				types:  c.StringSlice("type"),
				access: c.Bool("accessor"),
				index:  c.Int("enum"),
			}
			return g.generate()
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
	index  int
	buf    *bytes.Buffer
	pkg    *Package
	log    func(format string, args ...any)
}

func (g *Generator) generate() error {
	g.parse()
	return nil
}

func (g *Generator) parse() {
	pkg := fn.Panic1(packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(g.tags, " "))},
		Tests:      false,
		Logf:       g.log,
	}, g.files...))
	if len(pkg) != 1 {
		panic(fmt.Errorf("only one package for each generation.%d found with %s", len(pkg),
			strings.Join(g.files, ",")))
	}
	g.addPackage(pkg[0])
}

func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}
	var i int
	var f *ast.File
	for i, f = range pkg.Syntax {
		g.pkg.files[i] = &File{
			pkg:  g.pkg,
			file: f,
		}
	}
	g.lookupTypes()
}

func (g *Generator) lookupTypes() {

	var f *File
	for _, f = range g.files {
		if f.file != nil {
			ast.Inspect(f.file, f.genDecl)
		}
	}
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}
type File struct {
	pkg   *Package
	file  *ast.File
	types []*Type
}

func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		return true
	}
	var typeName string
	var spec ast.Spec
	var ts *ast.TypeSpec
	for _, spec = range decl.Specs {
		if ts, ok = spec.(*ast.TypeSpec); ok {
			typeName = ts.Name.Name
			f.types = append(f.types, &Type{
				typeName: typeName,
				fields:   nil,
				file:     f,
				spec:     ts,
			})
		}
	}
	return false
}

type Type struct {
	typeName string
	fields   []Field
	file     *File
	spec     *ast.TypeSpec
}
type Field struct {
	Name   string
	Column string
}
