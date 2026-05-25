package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Ident(t *testing.T) {
	got := typeString(&ast.Ident{Name: "string"})
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}

func TestTypeString_Star(t *testing.T) {
	got := typeString(&ast.StarExpr{X: &ast.Ident{Name: "int"}})
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}

func TestTypeString_Selector(t *testing.T) {
	got := typeString(&ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "NullString"}})
	if got != "sql.NullString" {
		t.Fatalf("expected sql.NullString, got %s", got)
	}
}

func TestTypeString_Array(t *testing.T) {
	got := typeString(&ast.ArrayType{Elt: &ast.Ident{Name: "byte"}})
	if got != "[]byte" {
		t.Fatalf("expected []byte, got %s", got)
	}
}

func TestTypeString_Interface(t *testing.T) {
	got := typeString(&ast.InterfaceType{Methods: &ast.FieldList{}})
	if got != "interface{}" {
		t.Fatalf("expected interface{}, got %s", got)
	}
}
