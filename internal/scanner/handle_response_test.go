//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_JSON 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestHandleResponse_JSON(t *testing.T) {
	ep := &Endpoint{}
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
