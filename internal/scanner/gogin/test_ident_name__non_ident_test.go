//ff:func feature=scan type=extract control=sequence
//ff:what TestIdentName_NonIdent 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIdentName_NonIdent(t *testing.T) {
	got := identName(&ast.BasicLit{Value: "42"})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
