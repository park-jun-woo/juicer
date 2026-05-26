//ff:func feature=sql type=test control=sequence
//ff:what TestCollectInlineSQLArgs_WithSQLCov 테스트
package sqls

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectInlineSQLArgs_WithSQLCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "ExecContext"}},
			Args: []ast.Expr{
				&ast.Ident{Name: "ctx"},
				&ast.BasicLit{Kind: token.STRING, Value: `"INSERT INTO users (name) VALUES ($1)"`},
			},
		}},
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "Close"}},
		}},
	}}
	collectInlineSQLArgs(body)
}
