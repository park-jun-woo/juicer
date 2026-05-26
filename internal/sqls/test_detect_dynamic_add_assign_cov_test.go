//ff:func feature=sql type=test control=sequence
//ff:what TestDetectDynamic_AddAssignCov 테스트
package sqls

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestDetectDynamic_AddAssignCov(t *testing.T) {
	body := &ast.BlockStmt{List: []ast.Stmt{
		&ast.AssignStmt{Tok: token.ADD_ASSIGN, Lhs: []ast.Expr{&ast.Ident{Name: "q"}}, Rhs: []ast.Expr{&ast.Ident{Name: "s"}}},
	}}
	if !detectDynamic(body) {
		t.Fatal("expected true for +=")
	}
}
