//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_RedirectWithStatus 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleResponse_RedirectWithStatus(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: `"/x"`},
		&ast.BasicLit{Kind: token.INT, Value: "301"},
	}}
	handleResponse(&ep, call, "redirect", nil, "handler")
	if ep.Responses[0].Status != "301" {
		t.Fatalf("redirect status = %s", ep.Responses[0].Status)
	}
}
