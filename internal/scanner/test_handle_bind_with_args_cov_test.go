//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBind_WithArgsCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestHandleBind_WithArgsCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	handleBind(ep, call, "ShouldBindJSON", info)
	if ep.Request.Body.VarName != "req" {
		t.Fatalf("expected req, got %s", ep.Request.Body.VarName)
	}
}
