//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
		},
	}

	handleQuery(&ep, call, "Query")

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.Query) != 1 || ep.Request.Query[0].Name != "page" {
		t.Fatalf("expected query param 'page', got %v", ep.Request.Query)
	}
}
