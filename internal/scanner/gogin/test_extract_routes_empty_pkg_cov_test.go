//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_EmptyPkgCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestExtractRoutes_EmptyPkgCov(t *testing.T) {
	fset := token.NewFileSet()
	file := &ast.File{
		Name:  &ast.Ident{Name: "main"},
		Decls: []ast.Decl{},
	}
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"main.go"},
		Fset:            fset,
	}
	result, _ := extractRoutes([]*packages.Package{pkg}, ".")
	if result == nil {
		result = []scanner.Endpoint{}
	}
}
