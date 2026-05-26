//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_DuplicateCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleForm_DuplicateCov(t *testing.T) {
	ep := &Endpoint{Request: &Request{FormFields: []Param{{Name: "name", Type: "string"}}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"name"`}}}
	handleForm(ep, call)
	if len(ep.Request.FormFields) != 1 {
		t.Fatal("expected 1 (deduped)")
	}
}
