package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExprString_Ident(t *testing.T) {
	got := exprString(&ast.Ident{Name: "x"})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}

func TestExprString_SelectorExpr(t *testing.T) {
	got := exprString(&ast.SelectorExpr{
		X:   &ast.Ident{Name: "pkg"},
		Sel: &ast.Ident{Name: "Func"},
	})
	if got != "pkg.Func" {
		t.Fatalf("expected pkg.Func, got %s", got)
	}
}

func TestExprString_StarExpr(t *testing.T) {
	got := exprString(&ast.StarExpr{X: &ast.Ident{Name: "int"}})
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}

func TestExprString_CompositeLit(t *testing.T) {
	got := exprString(&ast.CompositeLit{Type: &ast.Ident{Name: "Foo"}})
	if got != "Foo{}" {
		t.Fatalf("expected Foo{}, got %s", got)
	}
}

func TestExprString_CompositeLitNoType(t *testing.T) {
	got := exprString(&ast.CompositeLit{})
	if got != "{}" {
		t.Fatalf("expected {}, got %s", got)
	}
}

func TestExprString_UnaryExpr(t *testing.T) {
	got := exprString(&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "x"}})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}

func TestExprString_CallExpr(t *testing.T) {
	got := exprString(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}})
	if got != "f()" {
		t.Fatalf("expected f(), got %s", got)
	}
}

func TestExprString_ArrayType(t *testing.T) {
	got := exprString(&ast.ArrayType{Elt: &ast.Ident{Name: "int"}})
	if got != "[]int" {
		t.Fatalf("expected []int, got %s", got)
	}
}

func TestExprString_InterfaceType(t *testing.T) {
	got := exprString(&ast.InterfaceType{Methods: &ast.FieldList{}})
	if got != "interface{}" {
		t.Fatalf("expected interface{}, got %s", got)
	}
}

func TestExprString_BasicLit(t *testing.T) {
	got := exprString(&ast.BasicLit{Value: "42"})
	if got != "42" {
		t.Fatalf("expected 42, got %s", got)
	}
}

func TestExprString_MapType(t *testing.T) {
	got := exprString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}})
	if got != "map[string]int" {
		t.Fatalf("expected map[string]int, got %s", got)
	}
}

func TestExprString_IndexExpr(t *testing.T) {
	got := exprString(&ast.IndexExpr{X: &ast.Ident{Name: "arr"}, Index: &ast.Ident{Name: "i"}})
	if got != "arr[i]" {
		t.Fatalf("expected arr[i], got %s", got)
	}
}
