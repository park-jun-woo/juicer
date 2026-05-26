//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_GroupNoParentCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_GroupNoParentCov(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "g"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "unknown"}, Sel: &ast.Ident{Name: "Group"}}},
		},
	}
	processAssign(stmt, "gin", map[string]*routerInfo{})
}
