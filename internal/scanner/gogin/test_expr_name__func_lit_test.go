//ff:func feature=scan type=extract control=sequence
//ff:what TestExprName_FuncLit 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_FuncLit(t *testing.T) {
	got := exprName(&ast.FuncLit{
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	})
	if got != "(inline)" {
		t.Fatalf("expected (inline), got %s", got)
	}
}
