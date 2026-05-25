package scanner

import (
	"go/ast"
	"testing"
)

func TestExprName_Ident(t *testing.T) {
	got := exprName(&ast.Ident{Name: "handler"})
	if got != "handler" {
		t.Fatalf("expected handler, got %s", got)
	}
}

func TestExprName_SelectorExpr(t *testing.T) {
	got := exprName(&ast.SelectorExpr{
		X:   &ast.Ident{Name: "h"},
		Sel: &ast.Ident{Name: "Create"},
	})
	if got != "h.Create" {
		t.Fatalf("expected h.Create, got %s", got)
	}
}

func TestExprName_FuncLit(t *testing.T) {
	got := exprName(&ast.FuncLit{
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	})
	if got != "(inline)" {
		t.Fatalf("expected (inline), got %s", got)
	}
}

func TestExprName_CallExpr(t *testing.T) {
	got := exprName(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}})
	if got != "f()" {
		t.Fatalf("expected f(), got %s", got)
	}
}

func TestExprName_Nil(t *testing.T) {
	got := exprName(nil)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
