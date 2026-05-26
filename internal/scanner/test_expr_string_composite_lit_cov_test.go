//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_CompositeLitCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_CompositeLitCov(t *testing.T) {
	got := exprString(&ast.CompositeLit{Type: &ast.Ident{Name: "Foo"}})
	if got != "Foo{}" {
		t.Fatalf("expected Foo{}, got %s", got)
	}
	got = exprString(&ast.CompositeLit{})
	if got != "{}" {
		t.Fatalf("expected {}, got %s", got)
	}
}
