//ff:func feature=scan type=test control=sequence
//ff:what resolveGroupPrefixFunc 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
	"golang.org/x/tools/go/packages"
)

func TestResolveGroupPrefixFunc(t *testing.T) {
	fset := token.NewFileSet()
	info := &types.Info{}
	pkg := &packages.Package{
		Fset:      fset,
		TypesInfo: info,
	}
	fn := &ast.FuncDecl{
		Name: &ast.Ident{Name: "setup"},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}
	epIndex := map[struct{ file string; line int }]int{}
	resolveGroupPrefixFunc(fn, "gin", pkg, nil, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndex)
}
