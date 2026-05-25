package scanner

import (
	"go/ast"
	"testing"
)

func TestGinPkgName_DefaultImport(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/gin-gonic/gin"`}},
		},
	}
	got := ginPkgName(file)
	if got != "gin" {
		t.Fatalf("expected gin, got %s", got)
	}
}

func TestGinPkgName_AliasedImport(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Name: &ast.Ident{Name: "g"}, Path: &ast.BasicLit{Value: `"github.com/gin-gonic/gin"`}},
		},
	}
	got := ginPkgName(file)
	if got != "g" {
		t.Fatalf("expected g, got %s", got)
	}
}

func TestGinPkgName_NoGin(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"fmt"`}},
		},
	}
	got := ginPkgName(file)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
