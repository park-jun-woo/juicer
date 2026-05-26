//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_GroupCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestProcessAssign_GroupCov(t *testing.T) {
	routers := map[string]*routerInfo{"r": {prefix: "/api"}}
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "v1"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "r"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: 9, Value: `"/v1"`}},
			},
		},
	}
	processAssign(stmt, "gin", routers)
	if _, ok := routers["v1"]; !ok {
		t.Fatal("expected router v1")
	}
}
