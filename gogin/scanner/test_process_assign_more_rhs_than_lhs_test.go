//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_MoreRhsThanLhs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_MoreRhsThanLhs(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "a"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}}},
			&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "New"}}},
		},
	}
	routers := make(map[string]*routerInfo)
	processAssign(stmt, "gin", routers)
	if _, ok := routers["a"]; !ok {
		t.Error("expected 'a' router to be registered")
	}
}
