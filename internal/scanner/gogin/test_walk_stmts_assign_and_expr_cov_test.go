//ff:func feature=scan type=test control=sequence
//ff:what TestWalkStmts_AssignAndExprCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestWalkStmts_AssignAndExprCov(t *testing.T) {
	fset := token.NewFileSet()
	fset.AddFile("test.go", -1, 100)
	routers := map[string]*routerInfo{"r": {}}
	stmts := []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "v1"}},
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Group"}},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/v1"`}},
			}},
		},
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}, &ast.Ident{Name: "handler"}},
		}},
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Use"}},
			Args: []ast.Expr{&ast.Ident{Name: "mw"}},
		}},
		&ast.ExprStmt{X: &ast.Ident{Name: "skip"}},
		&ast.BlockStmt{List: nil},
	}
	var out []scanner.Endpoint
	walkStmts(stmts, "gin", "test.go", fset, routers, &out, map[int][]ast.Expr{})
}
