//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_UnknownCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprName_UnknownCov(t *testing.T) {
	got := exprName(&ast.CompositeLit{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
