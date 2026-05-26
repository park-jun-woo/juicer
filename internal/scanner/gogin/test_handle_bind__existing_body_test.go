//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleBind_ExistingBody 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleBind_ExistingBody(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{VarName: "existing"}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	handleBind(ep, call, "ShouldBindJSON", nil)
	if ep.Request.Body.VarName != "existing" {
		t.Fatal("should not overwrite existing body")
	}
}
