package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_Valid(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "H"}}
	if !isGinHSelector(sel) {
		t.Fatal("expected true")
	}
}

func TestIsGinHSelector_NonSelector(t *testing.T) {
	if isGinHSelector(&ast.Ident{Name: "x"}) {
		t.Fatal("expected false")
	}
}

func TestIsGinHSelector_NonGin(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "foo"}, Sel: &ast.Ident{Name: "H"}}
	if isGinHSelector(sel) {
		t.Fatal("expected false")
	}
}
