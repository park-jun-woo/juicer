//ff:func feature=sql type=test control=sequence
//ff:what TestDetectCRUD_QueryContextCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectCRUD_QueryContextCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "QueryContext"}},
		}},
	}}
	if detectCRUD(body) != "SELECT" {
		t.Fatal("expected SELECT")
	}
}
