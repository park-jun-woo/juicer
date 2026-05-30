//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr_NotFound 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestFindInfoForExpr_NotFound(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\nvar x = 1\n", 0)
	pkgs := []*packages.Package{{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}}}

	if got := findInfoForExpr(ast.NewIdent("synthetic"), pkgs); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
