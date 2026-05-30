//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_NoArgs 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleQuery_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handleQuery(&ep, &ast.CallExpr{}, "GET")
	if ep.Request != nil {
		t.Fatalf("expected no request, got %v", ep.Request)
	}
}
