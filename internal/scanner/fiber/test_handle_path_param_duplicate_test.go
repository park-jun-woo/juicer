//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_Duplicate 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandlePathParam_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}}}
	handlePathParam(&ep, call)
	handlePathParam(&ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param (dedup), got %d", len(ep.Request.PathParams))
	}
}
