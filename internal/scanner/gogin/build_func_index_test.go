//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_Nil 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_Nil(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil {
		t.Fatal("expected non-nil")
	}

	// pkg with nil TypesInfo
	pkg := &packages.Package{}
	idx = buildFuncIndex([]*packages.Package{pkg})
	if len(idx.byPos) != 0 {
		t.Fatal("expected empty")
	}

	// pkg with TypesInfo and syntax
	fset := token.NewFileSet()
	f := fset.AddFile("test.go", -1, 100)
	pos := f.Pos(1)
	fnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "Foo", NamePos: pos},
		Body: &ast.BlockStmt{},
		Type: &ast.FuncType{},
	}
	file := &ast.File{
		Decls: []ast.Decl{fnDecl, &ast.GenDecl{}},
	}
	pkg2 := &packages.Package{
		TypesInfo: &types.Info{},
		Syntax:    []*ast.File{file},
	}
	idx = buildFuncIndex([]*packages.Package{pkg2})
	if len(idx.byPos) != 1 {
		t.Fatalf("expected 1, got %d", len(idx.byPos))
	}
}
