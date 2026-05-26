//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_NonCall 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_NonCall(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.Ident{Name: "y"}},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
	if len(routers) != 0 {
		t.Fatal("expected no routers")
	}
}
