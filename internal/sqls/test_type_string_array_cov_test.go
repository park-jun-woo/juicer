//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_ArrayCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_ArrayCov(t *testing.T) {
	if typeString(&ast.ArrayType{Elt: &ast.Ident{Name: "int"}}) != "[]int" {
		t.Fatal("expected []int")
	}
}
