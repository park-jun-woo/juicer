//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_EmptyName 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandlePathParam_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}
	handlePathParam(&ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "v"}}})
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
}
