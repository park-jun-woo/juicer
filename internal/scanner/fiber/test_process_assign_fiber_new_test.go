//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_FiberNew 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_FiberNew(t *testing.T) {
	routers := map[string]*routerInfo{}
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "app"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "fiber"},
					Sel: &ast.Ident{Name: "New"},
				},
			},
		},
		Tok: token.DEFINE,
	}

	processAssign(stmt, "fiber", routers)
	if _, ok := routers["app"]; !ok {
		t.Fatal("expected app in routers")
	}
}
