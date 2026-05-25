//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_GinInit 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_GinInit(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "r"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "gin"},
					Sel: &ast.Ident{Name: "Default"},
				},
			},
		},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
	if _, ok := routers["r"]; !ok {
		t.Fatal("expected router r")
	}
}
