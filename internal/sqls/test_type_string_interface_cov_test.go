//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_InterfaceCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_InterfaceCov(t *testing.T) {
	if typeString(&ast.InterfaceType{Methods: &ast.FieldList{}}) != "interface{}" {
		t.Fatal("expected interface{}")
	}
}
