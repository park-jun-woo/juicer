package sqls

import (
	"go/ast"
	"testing"
)

func TestReceiverTypeName_Star(t *testing.T) {
	expr := &ast.StarExpr{X: &ast.Ident{Name: "UserRepo"}}
	got := receiverTypeName(expr)
	if got != "UserRepo" {
		t.Fatalf("expected UserRepo, got %s", got)
	}
}

func TestReceiverTypeName_Ident(t *testing.T) {
	got := receiverTypeName(&ast.Ident{Name: "UserRepo"})
	if got != "UserRepo" {
		t.Fatalf("expected UserRepo, got %s", got)
	}
}

func TestReceiverTypeName_Unknown(t *testing.T) {
	got := receiverTypeName(&ast.BasicLit{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
