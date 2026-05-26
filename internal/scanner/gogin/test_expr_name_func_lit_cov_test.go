//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_FuncLitCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_FuncLitCov(t *testing.T) {
	got := exprName(&ast.FuncLit{Type: &ast.FuncType{}, Body: &ast.BlockStmt{}})
	if got != "(inline)" {
		t.Fatalf("expected (inline), got %s", got)
	}
}
