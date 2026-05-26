//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr_NilPkgs 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr_NilPkgs(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	result := findInfoForExpr(expr, nil)
	if result != nil {
		t.Fatal("expected nil")
	}

	// pkg with nil TypesInfo
	pkg := &packages.Package{}
	result = findInfoForExpr(expr, []*packages.Package{pkg})
	if result != nil {
		t.Fatal("expected nil for nil TypesInfo")
	}

	// pkg with TypesInfo and matching file range
	fset := token.NewFileSet()
	f := fset.AddFile("test.go", 1, 100)
	// Create a file node that spans the full file range using a FuncDecl
	fnDecl := &ast.FuncDecl{
		Name: &ast.Ident{Name: "main", NamePos: f.Pos(5)},
		Type: &ast.FuncType{Func: f.Pos(1)},
		Body: &ast.BlockStmt{Lbrace: f.Pos(20), Rbrace: f.Pos(95)},
	}
	file := &ast.File{
		Package:  f.Pos(0),
		Name:     &ast.Ident{Name: "main", NamePos: f.Pos(1)},
		Decls:    []ast.Decl{fnDecl},
		FileEnd:  f.Pos(99),
	}
	ident := &ast.Ident{Name: "x", NamePos: f.Pos(50)}
	info := &types.Info{}
	pkg2 := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{file},
	}
	result = findInfoForExpr(ident, []*packages.Package{pkg2})
	if result != info {
		t.Fatal("expected matching info")
	}
}

