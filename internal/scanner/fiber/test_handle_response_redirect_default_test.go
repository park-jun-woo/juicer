//ff:func feature=scan type=test control=sequence
//ff:what TestHandleResponse_RedirectDefault 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleResponse_RedirectDefault(t *testing.T) {
	ep := scanner.Endpoint{}
	handleResponse(&ep, &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/x"`}}}, "redirect", nil, "handler")
	if ep.Responses[0].Status != "302" {
		t.Fatalf("redirect default status = %s", ep.Responses[0].Status)
	}
}
