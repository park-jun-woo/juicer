//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestFindInfoForExpr(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", "package m\nvar x = 1\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	info := &types.Info{}

	pkgs := []*packages.Package{
		{Syntax: []*ast.File{file}, TypesInfo: nil},
		{Syntax: []*ast.File{file}, TypesInfo: info},
	}

	var anExpr ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if e, ok := n.(*ast.BasicLit); ok && anExpr == nil {
			anExpr = e
			return false
		}
		return true
	})

	got := findInfoForExpr(anExpr, pkgs)
	if got != info {
		t.Fatalf("expected the non-nil TypesInfo, got %v", got)
	}
}
