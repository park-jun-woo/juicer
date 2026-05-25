//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_Group 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_Group(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "api"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "r"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}},
			},
		},
	}
	routers := map[string]*routerInfo{"r": {prefix: ""}}
	processAssign(stmt, "gin", routers)
	if _, ok := routers["api"]; !ok {
		t.Fatal("expected router api")
	}
}
