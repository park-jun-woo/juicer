//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_IndexOutOfRangeCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestExtractRoutes_IndexOutOfRangeCov(t *testing.T) {
	fset := token.NewFileSet()
	file1 := &ast.File{Name: &ast.Ident{Name: "main"}, Decls: []ast.Decl{}}
	file2 := &ast.File{Name: &ast.Ident{Name: "main"}, Decls: []ast.Decl{}}
	pkg := &packages.Package{
		Syntax:          []*ast.File{file1, file2},
		CompiledGoFiles: []string{"main.go"},
		Fset:            fset,
	}
	result, _ := extractRoutes([]*packages.Package{pkg}, ".")
	_ = result
}
