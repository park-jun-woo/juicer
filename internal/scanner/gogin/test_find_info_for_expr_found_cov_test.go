//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr_FoundCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr_FoundCov(t *testing.T) {
	fset := token.NewFileSet()
	f := fset.AddFile("test.go", -1, 100)
	f.SetLinesForContent(make([]byte, 100))

	base := f.Base()
	ident := &ast.Ident{Name: "x", NamePos: token.Pos(base + 10)}
	// Create a file with declarations spanning a range so Pos()..End() covers the ident
	fnIdent := &ast.Ident{Name: "foo", NamePos: token.Pos(base + 50)}
	fnDecl := &ast.FuncDecl{
		Name: fnIdent,
		Type: &ast.FuncType{Func: token.Pos(base + 45)},
		Body: &ast.BlockStmt{
			Lbrace: token.Pos(base + 60),
			Rbrace: token.Pos(base + 70),
		},
	}
	file := &ast.File{
		Name:    &ast.Ident{Name: "main", NamePos: token.Pos(base)},
		Package: token.Pos(base),
		Decls:   []ast.Decl{fnDecl},
	}
	info := &types.Info{}
	pkg := &packages.Package{
		TypesInfo: info,
		Syntax:    []*ast.File{file},
		Fset:      fset,
	}
	result := findInfoForExpr(ident, []*packages.Package{pkg})
	if result != info {
		t.Fatal("expected info")
	}
}
