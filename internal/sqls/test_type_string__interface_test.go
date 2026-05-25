//ff:func feature=sql type=parse control=sequence
//ff:what TestTypeString_Interface 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Interface(t *testing.T) {
	got := typeString(&ast.InterfaceType{Methods: &ast.FieldList{}})
	if got != "interface{}" {
		t.Fatalf("expected interface{}, got %s", got)
	}
}
