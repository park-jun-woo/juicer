//ff:func feature=scan type=test control=sequence
//ff:what TestStringLitValue_String 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringLitValue_String(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.STRING, Value: `"hello"`}
	got := stringLitValue(expr)
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}
