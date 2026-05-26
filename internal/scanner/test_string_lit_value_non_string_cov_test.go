//ff:func feature=scan type=test control=sequence
//ff:what TestStringLitValue_NonStringCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestStringLitValue_NonStringCov(t *testing.T) {
	got := stringLitValue(&ast.Ident{Name: "x"})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
