//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_NonIdentLhsCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_NonIdentLhsCov(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "a"}, Sel: &ast.Ident{Name: "b"}}},
		Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}}}},
	}
	processAssign(stmt, "gin", map[string]*routerInfo{})
}
