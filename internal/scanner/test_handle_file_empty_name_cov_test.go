//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_EmptyNameCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleFile_EmptyNameCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "varName"}}}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected no request for non-string arg")
	}
}
