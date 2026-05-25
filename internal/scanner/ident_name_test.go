package scanner

import (
	"go/ast"
	"testing"
)

func TestIdentName_Ident(t *testing.T) {
	got := identName(&ast.Ident{Name: "x"})
	if got != "x" {
		t.Fatalf("expected x, got %s", got)
	}
}

func TestIdentName_NonIdent(t *testing.T) {
	got := identName(&ast.BasicLit{Value: "42"})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
