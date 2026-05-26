//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUDFromAST_InsertCov 테스트
package sqls

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestRefineCRUDFromAST_InsertCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "ExecContext"}},
			Args: []ast.Expr{&ast.Ident{Name: "ctx"}, &ast.BasicLit{Kind: token.STRING, Value: `"INSERT INTO users"`}},
		}},
	}}
	if refineCRUDFromAST(body) != "INSERT" {
		t.Fatal("expected INSERT")
	}
}
