//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_GenGoIncluded 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_GenGoIncluded(t *testing.T) {
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
	if len(eps) != 1 {
		t.Errorf("expected 1 route from .gen.go, got %d", len(eps))
	}
}
