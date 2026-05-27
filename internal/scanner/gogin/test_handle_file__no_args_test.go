//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleFile_NoArgs 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleFile_NoArgs(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected nil request")
	}
}
