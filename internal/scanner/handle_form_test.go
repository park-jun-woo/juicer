//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_Basic 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleForm_Basic(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"name"`}},
	}
	handleForm(ep, call)
	if len(ep.Request.FormFields) != 1 {
		t.Fatal("expected 1 form field")
	}
}
