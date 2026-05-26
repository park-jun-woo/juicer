//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_ArrayTypeCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_ArrayTypeCov(t *testing.T) {
	got := exprString(&ast.ArrayType{Elt: &ast.Ident{Name: "byte"}})
	if got != "[]byte" {
		t.Fatalf("expected []byte, got %s", got)
	}
}
