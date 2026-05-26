//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_CompositeLitNoType 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_CompositeLitNoType(t *testing.T) {
	got := exprString(&ast.CompositeLit{})
	if got != "{}" {
		t.Fatalf("expected {}, got %s", got)
	}
}
