//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_JSON 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_JSON(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "200"},
			&ast.Ident{Name: "data"},
		},
	}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	handleResponse(ep, call, "json", info, "handler")
	if len(ep.Responses) != 1 {
		t.Fatal("expected 1 response")
	}
}

