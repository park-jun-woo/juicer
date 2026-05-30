//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_EmptyName 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestHandleFile_EmptyName(t *testing.T) {
	ep := scanner.Endpoint{}

	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "nameVar"}}}
	handleFile(&ep, call)
	if ep.Request != nil {
		t.Fatalf("expected no request for empty name, got %v", ep.Request)
	}
}
