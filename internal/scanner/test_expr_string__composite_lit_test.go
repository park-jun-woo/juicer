//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_CompositeLit 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_CompositeLit(t *testing.T) {
	got := exprString(&ast.CompositeLit{Type: &ast.Ident{Name: "Foo"}})
	if got != "Foo{}" {
		t.Fatalf("expected Foo{}, got %s", got)
	}
}
