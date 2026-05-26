//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_EllipsisCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_EllipsisCov(t *testing.T) {
	if typeString(&ast.Ellipsis{Elt: &ast.Ident{Name: "int"}}) != "...int" {
		t.Fatal("expected ...int")
	}
}
