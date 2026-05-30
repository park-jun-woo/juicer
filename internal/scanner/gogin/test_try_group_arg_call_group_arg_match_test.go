//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_GroupArgMatch 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryGroupArgCall_GroupArgMatch(t *testing.T) {

	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "register"},
		Args: []ast.Expr{&ast.Ident{Name: "api"}},
	}
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"api": {prefix: "/api"}},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    goginEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}
