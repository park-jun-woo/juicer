//ff:func feature=sql type=test control=sequence
//ff:what TestCollectSQLFragments_WithSQLCov 테스트
package sqls

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectSQLFragments_WithSQLCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.BasicLit{Kind: token.STRING, Value: "`SELECT id, name FROM users WHERE id = $1`"}},
		&ast.ExprStmt{X: &ast.BasicLit{Kind: token.STRING, Value: `"not a backtick"`}},
		&ast.ExprStmt{X: &ast.BasicLit{Kind: token.STRING, Value: "`short`"}},
		&ast.ExprStmt{X: &ast.BasicLit{Kind: token.STRING, Value: "`this is not sql at all here`"}},
	}}
	result := collectSQLFragments(body)
	if len(result) < 1 {
		t.Fatal("expected at least 1 SQL fragment")
	}
}
