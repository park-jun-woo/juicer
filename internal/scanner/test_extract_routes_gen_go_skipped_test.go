//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractRoutes_GenGoSkipped 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_GenGoSkipped(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.GET("/test", handler)
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.gen.go", src, 0)

	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"main.gen.go"},
		Fset:            fset,
	}

	eps, _ := extractRoutes([]*packages.Package{pkg}, ".")
	if len(eps) != 0 {
		t.Errorf("expected 0 for .gen.go, got %d", len(eps))
	}
}
