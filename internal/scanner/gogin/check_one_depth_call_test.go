//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_NilInfoOld 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestCheckOneDepthCall_NilInfoOld(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", nil, idx)

	// no ctx arg
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	call2 := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "other"}},
	}
	checkOneDepthCall(ep, call2, "c", info, idx)

	// with ctx arg but no valid target
	call3 := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "f"},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	checkOneDepthCall(ep, call3, "c", info, idx)
}


func TestCheckOneDepthCall_HasCtxUnresolvedZ(t *testing.T) {
	ep := &scanner.Endpoint{}
	// passes "c" -> hasCtx true; empty info -> resolveCallTarget invalid -> return
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
