//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_MapTypeCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_MapTypeCov(t *testing.T) {
	got := exprString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}})
	if got != "map[string]int" {
		t.Fatalf("expected map[string]int, got %s", got)
	}
}
