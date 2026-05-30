//ff:func feature=scan type=test control=iteration dimension=1
//ff:what findInfoForExpr — 표현 소속 TypesInfo 검색 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", "package m\nvar x = 1\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	info := &types.Info{}

	// package with nil TypesInfo -> skipped; second package has it
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

func TestFindInfoForExpr_NotFound(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\nvar x = 1\n", 0)
	pkgs := []*packages.Package{{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}}}
	// an expr with no position (synthetic) -> not within any file range
	if got := findInfoForExpr(ast.NewIdent("synthetic"), pkgs); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
