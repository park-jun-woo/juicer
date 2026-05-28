//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_Group 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_Group(t *testing.T) {
	routers := map[string]*routerInfo{
		"app": {prefix: ""},
	}
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "api"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "app"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
				},
			},
		},
		Tok: token.DEFINE,
	}

	processAssign(stmt, "fiber", routers)
	ri, ok := routers["api"]
	if !ok {
		t.Fatal("expected api in routers")
	}
	if ri.prefix != "/api" {
		t.Fatalf("expected /api, got %s", ri.prefix)
	}
}
