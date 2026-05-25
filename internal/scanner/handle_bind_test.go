package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestHandleBind_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	info := &types.Info{}
	handleBind(ep, call, "ShouldBindJSON", info)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body to be set")
	}
}

func TestHandleBind_ExistingBody(t *testing.T) {
	ep := &Endpoint{Request: &Request{Body: &Body{VarName: "existing"}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	handleBind(ep, call, "ShouldBindJSON", nil)
	if ep.Request.Body.VarName != "existing" {
		t.Fatal("should not overwrite existing body")
	}
}
