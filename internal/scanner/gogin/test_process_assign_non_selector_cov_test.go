//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_NonSelectorCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_NonSelectorCov(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.Ident{Name: "f"}}},
	}
	processAssign(stmt, "gin", map[string]*routerInfo{})
}
