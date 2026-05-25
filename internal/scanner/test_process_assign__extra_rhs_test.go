//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_ExtraRhs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_ExtraRhs(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.Ident{Name: "a"}, &ast.Ident{Name: "b"}},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
}
