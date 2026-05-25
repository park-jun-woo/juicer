//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleBind_ExistingBody 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleBind_ExistingBody(t *testing.T) {
	ep := &Endpoint{Request: &Request{Body: &Body{VarName: "existing"}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	handleBind(ep, call, "ShouldBindJSON", nil)
	if ep.Request.Body.VarName != "existing" {
		t.Fatal("should not overwrite existing body")
	}
}
