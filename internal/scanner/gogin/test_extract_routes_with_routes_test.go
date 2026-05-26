//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractRoutes_WithRoutes 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_WithRoutes(t *testing.T) {
	src := `package main

import "github.com/gin-gonic/gin"

func setup() {
	r := gin.Default()
	r.POST("/b", handler)
	r.GET("/a", handler)
	r.GET("/b", handler)
}

func handler(c *gin.Context) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "main.go", src, 0)

	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"main.go"},
		Fset:            fset,
	}

	eps, _ := extractRoutes([]*packages.Package{pkg}, ".")
	if len(eps) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(eps))
	}
	// Should be sorted: /a GET, /b GET, /b POST
	if eps[0].Path != "/a" {
		t.Errorf("expected first path /a, got %q", eps[0].Path)
	}
	if eps[1].Path != "/b" && eps[1].Method != "GET" {
		t.Errorf("expected second /b GET")
	}
}
