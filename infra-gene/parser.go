package main

import (
	"bytes"
	"fmt"
	"github.com/ZenLiuCN/fn"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

type Tags map[string][]string

func NewTags(s string) Tags {
	if s == "" {
		return nil
	}
	t := make(Tags)
	for _, seg := range strings.Split(s, " ") {
		v := strings.Split(seg, ":")
		if len(v) == 1 {
			t[v[0]] = nil
		} else {
			val := strings.Split(v[1][1:len(v[1])-1], ",")
			for _, s2 := range val {
				ss := strings.TrimSpace(s2)
				if len(ss) > 0 {
					t[v[0]] = append(t[v[0]], ss)
				}
			}
		}
	}
	return t
}
func NewTagsOf(s *ast.BasicLit) Tags {
	if s == nil {
		return nil
	}
	return NewTags(s.Value[1 : len(s.Value)-1])
}

type Context struct {
	Pkg  []*Package
	Logf func(format string, args ...any)
}

func (c *Context) Printf(format string, args ...any) {
	if c.Logf != nil {
		c.Logf(format, args...)
	}
}
func (c *Context) Parse(tags []string, files []string) {
	pkg := fn.Panic1(packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
		Tests:      false,
		Logf:       c.Logf,
	}, files...))
	c.Pkg = make([]*Package, len(pkg))
	for i, p := range pkg {
		c.Pkg[i] = &Package{
			Context: c,
			Name:    p.Name,
			Pkg:     p,
		}
		c.Pkg[i].ResolveFiles()
	}
}

type Package struct {
	*Context
	Name  string
	Pkg   *packages.Package
	Files []*File
}

func (p *Package) ResolveFiles() {
	p.Files = make([]*File, len(p.Pkg.Syntax))
	for i, f := range p.Pkg.Syntax {
		p.Files[i] = &File{
			Pkg:  p,
			file: f,
		}
	}
}

type File struct {
	Pkg  *Package
	file *ast.File
}

type Writer struct {
	Package string
	Imports fn.HashSet[string]
	buf     *bytes.Buffer
}

func WriterOf(buf *bytes.Buffer) *Writer {
	return &Writer{buf: buf}
}
func NewWriter() *Writer {
	return &Writer{buf: new(bytes.Buffer)}
}
func (s *Writer) Buffer() *bytes.Buffer {
	return s.buf
}
func (s *Writer) F(format string, args ...any) *Writer {
	if s.buf == nil {
		s.buf = new(bytes.Buffer)
	}
	_, _ = fmt.Fprintf(s.buf, format, args...)
	return s
}
func (s *Writer) Append(w *Writer) *Writer {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	for s2 := range w.Imports {
		s.Imports.Put(s2)
	}
	s.buf.Write(w.buf.Bytes())
	return s
}

func (s *Writer) Import(path string) *Writer {
	if s.Imports == nil {
		s.Imports = fn.NewHashSet[string]()
	}
	s.Imports.Put(path)
	return s
}
