//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_NoArgsCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleFile_NoArgsCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected no request")
	}
}
