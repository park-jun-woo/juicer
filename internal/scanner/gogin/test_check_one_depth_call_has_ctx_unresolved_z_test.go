//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_HasCtxUnresolvedZ 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_HasCtxUnresolvedZ(t *testing.T) {
	ep := &scanner.Endpoint{}

	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "respond"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}, &ast.BasicLit{Kind: token.INT, Value: "200"}},
	}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	info := &types.Info{Uses: map[*ast.Ident]types.Object{}, Selections: map[*ast.SelectorExpr]*types.Selection{}}
	checkOneDepthCall(ep, call, "c", info, idx)
	if len(ep.Responses) != 0 {
		t.Fatalf("expected no responses, got %d", len(ep.Responses))
	}
}
