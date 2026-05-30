//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_NoArgs 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleFile_NoArgs(t *testing.T) {
	ep := scanner.Endpoint{}
	handleFile(&ep, &ast.CallExpr{})
	if ep.Request != nil {
		t.Fatalf("expected no request for no args, got %v", ep.Request)
	}
}
