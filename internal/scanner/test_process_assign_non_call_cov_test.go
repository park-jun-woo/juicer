//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_NonCallCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_NonCallCov(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.Ident{Name: "y"}},
	}
	processAssign(stmt, "gin", map[string]*routerInfo{})
}
