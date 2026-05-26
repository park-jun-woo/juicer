//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_NoArgsCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleFile_NoArgsCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected no request")
	}
}
