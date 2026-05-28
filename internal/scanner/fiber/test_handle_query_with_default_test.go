//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_WithDefault 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_WithDefault(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
			&ast.BasicLit{Kind: token.STRING, Value: `"1"`},
		},
	}

	handleQuery(&ep, call, "Query")

	if ep.Request == nil || len(ep.Request.Query) != 1 {
		t.Fatal("expected 1 query param")
	}
	if ep.Request.Query[0].Default != "1" {
		t.Fatalf("expected default '1', got %s", ep.Request.Query[0].Default)
	}
}
