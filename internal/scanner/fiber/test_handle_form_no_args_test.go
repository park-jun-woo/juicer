//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_NoArgs 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleForm_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handleForm(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request for no args, got %v", ep.Request)
	}
}
