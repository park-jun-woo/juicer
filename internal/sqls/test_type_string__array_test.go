//ff:func feature=sql type=parse control=sequence
//ff:what TestTypeString_Array 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Array(t *testing.T) {
	got := typeString(&ast.ArrayType{Elt: &ast.Ident{Name: "byte"}})
	if got != "[]byte" {
		t.Fatalf("expected []byte, got %s", got)
	}
}
