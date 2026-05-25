//ff:func feature=scan type=extract control=sequence
//ff:what TestExprString_InterfaceType 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_InterfaceType(t *testing.T) {
	got := exprString(&ast.InterfaceType{Methods: &ast.FieldList{}})
	if got != "interface{}" {
		t.Fatalf("expected interface{}, got %s", got)
	}
}
