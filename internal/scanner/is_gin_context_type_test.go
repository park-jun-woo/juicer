package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_Valid(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "Context"},
		},
	}
	if !isGinContextType(expr) {
		t.Fatal("expected true")
	}
}

func TestIsGinContextType_NotStar(t *testing.T) {
	if isGinContextType(&ast.Ident{Name: "int"}) {
		t.Fatal("expected false")
	}
}

func TestIsGinContextType_WrongName(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Request"},
		},
	}
	if isGinContextType(expr) {
		t.Fatal("expected false")
	}
}
