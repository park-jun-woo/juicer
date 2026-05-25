//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_MapType 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_MapType(t *testing.T) {
	got := exprString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}})
	if got != "map[string]int" {
		t.Fatalf("expected map[string]int, got %s", got)
	}
}
