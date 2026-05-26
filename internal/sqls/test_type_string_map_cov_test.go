//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_MapCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_MapCov(t *testing.T) {
	if typeString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}}) != "map[string]int" {
		t.Fatal("expected map[string]int")
	}
}
