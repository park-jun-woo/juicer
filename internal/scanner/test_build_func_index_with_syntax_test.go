//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_WithSyntax 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_WithSyntax(t *testing.T) {
	fset := token.NewFileSet()
	src := `package main
func hello() {}
`
	f := fset.AddFile("test.go", -1, len(src))
	f.SetLinesForContent([]byte(src))

	fnIdent := &ast.Ident{Name: "hello", NamePos: token.Pos(f.Base() + 14)}
	fnDecl := &ast.FuncDecl{
		Name: fnIdent,
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}
	file := &ast.File{
		Name:  &ast.Ident{Name: "main"},
		Decls: []ast.Decl{fnDecl},
	}
	info := &types.Info{}
	pkg := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{file},
	}
	idx := buildFuncIndex([]*packages.Package{pkg})
	if len(idx.byPos) != 1 {
		t.Fatalf("expected 1 func in index, got %d", len(idx.byPos))
	}
}
