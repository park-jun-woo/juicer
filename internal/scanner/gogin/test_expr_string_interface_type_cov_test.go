//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_InterfaceTypeCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_InterfaceTypeCov(t *testing.T) {
	got := exprString(&ast.InterfaceType{Methods: &ast.FieldList{}})
	if got != "interface{}" {
		t.Fatalf("expected interface{}, got %s", got)
	}
}
