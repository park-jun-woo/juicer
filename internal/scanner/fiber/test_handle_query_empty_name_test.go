//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_EmptyName 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleQuery_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}
	handleQuery(&ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "v"}}}, "GET")
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
	_ = token.STRING
}
