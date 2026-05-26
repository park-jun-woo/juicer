//ff:func feature=scan type=extract control=sequence
//ff:what TestCheckOneDepthCall_DefaultCase 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestCheckOneDepthCall_DefaultCase(t *testing.T) {
	ep := &scanner.Endpoint{}
	// Use a CallExpr with Fun as something that's not Ident or SelectorExpr
	call := &ast.CallExpr{
		Fun:  &ast.ParenExpr{X: &ast.Ident{Name: "someFunc"}},
		Args: []ast.Expr{&ast.Ident{Name: "c"}},
	}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	idx := &funcIndex{
		byPos: make(map[token.Pos]*ast.FuncDecl),
		info:  make(map[token.Pos]*types.Info),
	}
	checkOneDepthCall(ep, call, "c", info, idx)
	// Should return in default case
}
