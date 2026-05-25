//ff:func feature=scan type=extract control=sequence
//ff:what TestStringLitValue_NonBasicLit 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestStringLitValue_NonBasicLit(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	got := stringLitValue(expr)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
