package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinInit_Default(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Default"}}
	if !isGinInit(sel, "gin") {
		t.Fatal("expected true")
	}
}

func TestIsGinInit_New(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "New"}}
	if !isGinInit(sel, "gin") {
		t.Fatal("expected true")
	}
}

func TestIsGinInit_NotGin(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "http"}, Sel: &ast.Ident{Name: "Default"}}
	if isGinInit(sel, "gin") {
		t.Fatal("expected false")
	}
}

func TestIsGinInit_NonIdent(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.BasicLit{Value: "x"}, Sel: &ast.Ident{Name: "Default"}}
	if isGinInit(sel, "gin") {
		t.Fatal("expected false")
	}
}
