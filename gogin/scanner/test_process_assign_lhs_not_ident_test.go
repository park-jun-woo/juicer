//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_LhsNotIdent 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_LhsNotIdent(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "obj"}, Sel: &ast.Ident{Name: "field"}}},
		Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}}}},
	}
	routers := make(map[string]*routerInfo)
	processAssign(stmt, "gin", routers)
	if len(routers) != 0 {
		t.Error("expected 0 routers")
	}
}
