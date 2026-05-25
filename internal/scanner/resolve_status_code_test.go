package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveStatusCode_IntLit(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.INT, Value: "200"}
	got := resolveStatusCode(expr, nil)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}

func TestResolveStatusCode_NilInfo(t *testing.T) {
	expr := &ast.Ident{Name: "StatusOK"}
	got := resolveStatusCode(expr, nil)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}

func TestResolveStatusCode_NoMatch(t *testing.T) {
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	expr := &ast.Ident{Name: "x"}
	got := resolveStatusCode(expr, info)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}
