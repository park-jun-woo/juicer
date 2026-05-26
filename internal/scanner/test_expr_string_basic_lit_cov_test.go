//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_BasicLitCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprString_BasicLitCov(t *testing.T) {
	got := exprString(&ast.BasicLit{Kind: token.INT, Value: "42"})
	if got != "42" {
		t.Fatalf("expected 42, got %s", got)
	}
}
