//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_BasicLit 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_BasicLit(t *testing.T) {
	got := exprString(&ast.BasicLit{Value: "42"})
	if got != "42" {
		t.Fatalf("expected 42, got %s", got)
	}
}
