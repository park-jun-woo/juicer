//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_CallExprCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_CallExprCov(t *testing.T) {
	got := exprName(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}})
	if got != "f()" {
		t.Fatalf("expected f(), got %s", got)
	}
}
