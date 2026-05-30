//ff:func feature=scan type=test control=sequence
//ff:what stringLitValue — 문자열 리터럴 값 추출 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringLitValue(t *testing.T) {
	if got := stringLitValue(&ast.BasicLit{Kind: token.STRING, Value: `"hello"`}); got != "hello" {
		t.Errorf("string lit: %q", got)
	}
	// non-string lit
	if got := stringLitValue(&ast.BasicLit{Kind: token.INT, Value: "5"}); got != "" {
		t.Errorf("int lit should be empty, got %q", got)
	}
	// non-lit
	if got := stringLitValue(&ast.Ident{Name: "x"}); got != "" {
		t.Errorf("ident should be empty, got %q", got)
	}
}
