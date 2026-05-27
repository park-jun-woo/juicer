//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleResponse_Status 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_Status(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "204"}}}
	handleResponse(ep, call, "status", nil, "handler")
	if len(ep.Responses) != 1 {
		t.Fatal("expected 1 response")
	}
}
