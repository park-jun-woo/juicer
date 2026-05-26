//ff:func feature=scan type=test control=sequence
//ff:what TestIdentName_NonIdentCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIdentName_NonIdentCov(t *testing.T) {
	got := identName(&ast.CompositeLit{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
