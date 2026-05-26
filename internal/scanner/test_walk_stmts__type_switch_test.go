//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_TypeSwitchStmt — typeSwitch 내부 라우트 감지 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkStmts_TypeSwitchStmt(t *testing.T) {
	routeCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "r"},
			Sel: &ast.Ident{Name: "GET"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"/in-typeswitch"`},
			&ast.Ident{Name: "handler"},
		},
	}
	typeSwitchStmt := &ast.TypeSwitchStmt{
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.CaseClause{
					Body: []ast.Stmt{
						&ast.ExprStmt{X: routeCall},
					},
				},
			},
		},
	}

	routers := map[string]*routerInfo{"r": {}}
	var out []Endpoint
	fset := token.NewFileSet()
	walkStmts([]ast.Stmt{typeSwitchStmt}, "gin", "test.go", fset, routers, &out)
	if len(out) != 1 {
		t.Fatalf("expected 1 endpoint from typeswitch, got %d", len(out))
	}
	if out[0].Path != "/in-typeswitch" {
		t.Errorf("expected /in-typeswitch, got %s", out[0].Path)
	}
}
