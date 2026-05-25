//ff:func feature=scan type=extract control=sequence
//ff:what TestProcessAssign_GroupFromUnknown 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_GroupFromUnknown(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "g"}},
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "unknown"},
				Sel: &ast.Ident{Name: "Group"},
			},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}},
		}},
	}
	routers := make(map[string]*routerInfo)
	processAssign(stmt, "gin", routers)
	if len(routers) != 0 {
		t.Error("expected 0 routers for unknown parent")
	}
}
