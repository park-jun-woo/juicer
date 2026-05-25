//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleResponse_Status 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleResponse_Status(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "204"}}}
	handleResponse(ep, call, "status", nil, "handler")
	if len(ep.Responses) != 1 {
		t.Fatal("expected 1 response")
	}
}
