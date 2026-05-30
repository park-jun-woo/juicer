//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_MismatchedLhsRhs 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_MismatchedLhsRhs(t *testing.T) {

	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "a"}},
		Rhs: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "1"},
			&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "New"}}},
		},
		Tok: token.DEFINE,
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "fiber", routers)

	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}
