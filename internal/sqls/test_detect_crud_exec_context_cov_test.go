//ff:func feature=sql type=test control=sequence
//ff:what TestDetectCRUD_ExecContextCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestDetectCRUD_ExecContextCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "ExecContext"}},
		}},
	}}
	if detectCRUD(body) != "EXEC" {
		t.Fatal("expected EXEC")
	}
}
