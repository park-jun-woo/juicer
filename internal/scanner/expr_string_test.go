//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_Ident 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_Ident(t *testing.T) {
	got := exprString(&ast.Ident{Name: "x"})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}

