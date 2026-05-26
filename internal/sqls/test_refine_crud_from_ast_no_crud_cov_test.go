//ff:func feature=sql type=test control=sequence
//ff:what TestRefineCRUDFromAST_NoCRUDCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestRefineCRUDFromAST_NoCRUDCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "db"}, Sel: &ast.Ident{Name: "ExecContext"}},
		}},
	}}
	if refineCRUDFromAST(body) != "EXEC" {
		t.Fatal("expected EXEC")
	}
}
