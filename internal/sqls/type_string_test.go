//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_Ident 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Ident(t *testing.T) {
	if typeString(&ast.Ident{Name: "string"}) != "string" {
		t.Fatal("ident")
	}
	if typeString(&ast.StarExpr{X: &ast.Ident{Name: "int"}}) != "*int" {
		t.Fatal("star")
	}
	if typeString(&ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "DB"}}) != "sql.DB" {
		t.Fatal("selector")
	}
	if typeString(&ast.ArrayType{Elt: &ast.Ident{Name: "string"}}) != "[]string" {
		t.Fatal("array")
	}
	if typeString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}}) != "map[string]int" {
		t.Fatal("map")
	}
	if typeString(&ast.InterfaceType{}) != "interface{}" {
		t.Fatal("interface")
	}
	if typeString(&ast.Ellipsis{Elt: &ast.Ident{Name: "string"}}) != "...string" {
		t.Fatal("ellipsis")
	}
}

