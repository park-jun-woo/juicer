//ff:func feature=scan type=test control=sequence
//ff:what TestHandleFile_EmptyNameCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleFile_EmptyNameCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "varName"}}}
	handleFile(ep, call)
	if ep.Request != nil {
		t.Fatal("expected no request for non-string arg")
	}
}
