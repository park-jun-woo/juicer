package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestCollectStringParts_BasicLit(t *testing.T) {
	var parts []string
	lit := &ast.BasicLit{Kind: token.STRING, Value: `"/api"`}
	collectStringParts(lit, &parts)
	if len(parts) != 1 || parts[0] != "/api" {
		t.Fatalf("got %v", parts)
	}
}

func TestCollectStringParts_Binary(t *testing.T) {
	var parts []string
	expr := &ast.BinaryExpr{
		Op: token.ADD,
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/health"`},
	}
	collectStringParts(expr, &parts)
	if len(parts) != 2 {
		t.Fatalf("expected 2, got %d", len(parts))
	}
}
