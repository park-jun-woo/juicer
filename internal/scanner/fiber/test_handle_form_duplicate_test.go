//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_Duplicate 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleForm_Duplicate(t *testing.T) {
	ep := scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"name"`}}}
	handleForm(&ep, call)
	handleForm(&ep, call)
	if len(ep.Request.FormFields) != 1 {
		t.Fatalf("expected 1 form field (dedup), got %d", len(ep.Request.FormFields))
	}
}
