//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleResponse_File 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleResponse_File(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"file.txt"`}}}
	handleResponse(ep, call, "file", nil, "handler")
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "200" {
		t.Fatal("expected 200 for file response")
	}
}
