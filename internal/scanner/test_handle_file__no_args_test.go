//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleFile_NoArgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleFile_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected nil request")
	}
}
