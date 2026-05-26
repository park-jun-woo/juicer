//ff:func feature=scan type=extract control=sequence
//ff:what TestIdentName 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIdentName(t *testing.T) {
	if identName(&ast.Ident{Name: "foo"}) != "foo" {
		t.Error("expected 'foo'")
	}
	if identName(&ast.CompositeLit{}) != "" {
		t.Error("expected empty")
	}
}
