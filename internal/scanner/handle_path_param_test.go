//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_Basic 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandlePathParam_Basic(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}},
	}
	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatal("expected 1 path param")
	}
}

