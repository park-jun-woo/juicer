//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_NonSelectorFun 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_NonSelectorFun(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "a"}},
		Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.Ident{Name: "something"}}},
	}
	routers := make(map[string]*routerInfo)
	processAssign(stmt, "gin", routers)
	if len(routers) != 0 {
		t.Error("expected 0 routers")
	}
}
