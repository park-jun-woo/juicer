//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_ArrayType 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_ArrayType(t *testing.T) {
	got := exprString(&ast.ArrayType{Elt: &ast.Ident{Name: "int"}})
	if got != "[]int" {
		t.Fatalf("expected []int, got %s", got)
	}
}
