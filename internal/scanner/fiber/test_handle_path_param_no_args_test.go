//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_NoArgs 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandlePathParam_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handlePathParam(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
}
