package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString_BasicLit(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.STRING, Value: `"/api/v1"`}
	path, ok := extractPathString(expr)
	if !ok || path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s ok=%v", path, ok)
	}
}

func TestExtractPathString_BinaryExpr(t *testing.T) {
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
	}
	path, ok := extractPathString(expr)
	if !ok || path != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s ok=%v", path, ok)
	}
}

func TestExtractPathString_Unknown(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	_, ok := extractPathString(expr)
	if ok {
		t.Fatal("expected not ok")
	}
}
