//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandlePathParam_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"id"`},
		},
	}

	handlePathParam(&ep, call)

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("expected path param 'id', got %v", ep.Request.PathParams)
	}
}
