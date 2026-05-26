//ff:func feature=scan type=test control=sequence
//ff:what TestIdentName_Ident 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIdentName_Ident(t *testing.T) {
	got := identName(&ast.Ident{Name: "x"})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}

