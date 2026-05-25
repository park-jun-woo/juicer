package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_StarExpr(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "Engine"},
		},
	}
	if !isGinRouterType(expr, "gin") {
		t.Fatal("expected true")
	}
}

func TestIsGinRouterType_NotGin(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Server"},
		},
	}
	if isGinRouterType(expr, "gin") {
		t.Fatal("expected false")
	}
}

func TestIsGinRouterType_NotSelector(t *testing.T) {
	if isGinRouterType(&ast.Ident{Name: "x"}, "gin") {
		t.Fatal("expected false for non-selector")
	}
}
