package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringLitValue_String(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.STRING, Value: `"hello"`}
	got := stringLitValue(expr)
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

func TestStringLitValue_NonString(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.INT, Value: "42"}
	got := stringLitValue(expr)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}

func TestStringLitValue_NonBasicLit(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	got := stringLitValue(expr)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
