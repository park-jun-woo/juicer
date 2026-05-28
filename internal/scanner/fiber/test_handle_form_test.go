//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_Basic 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleForm_Basic(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"name"`},
		},
	}

	handleForm(&ep, call)

	if ep.Request == nil {
		t.Fatal("expected non-nil Request")
	}
	if len(ep.Request.FormFields) != 1 || ep.Request.FormFields[0].Name != "name" {
		t.Fatalf("expected form field 'name', got %v", ep.Request.FormFields)
	}
}
