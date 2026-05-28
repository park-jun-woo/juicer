//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_Duplicate 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{
		Request: &scanner.Request{
			Query: []scanner.Param{{Name: "page", Type: "string"}},
		},
	}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
		},
	}

	handleQuery(&ep, call, "Query")

	if len(ep.Request.Query) != 1 {
		t.Fatalf("expected 1 (no duplicate), got %d", len(ep.Request.Query))
	}
}
