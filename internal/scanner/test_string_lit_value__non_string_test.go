//ff:func feature=scan type=extract control=sequence
//ff:what TestStringLitValue_NonString 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringLitValue_NonString(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.INT, Value: "42"}
	got := stringLitValue(expr)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
