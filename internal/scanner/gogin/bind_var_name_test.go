//ff:func feature=scan type=test control=sequence
//ff:what TestBindVarName_UnaryAnd 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestBindVarName_UnaryAnd(t *testing.T) {
	expr := &ast.UnaryExpr{
		Op: token.AND,
		X:  &ast.Ident{Name: "req"},
	}
	got := bindVarName(expr)
	if got != "req" {
		t.Fatalf("got %q", got)
	}

	// non-unary expr
	got = bindVarName(&ast.Ident{Name: "data"})
	if got != "data" {
		t.Fatalf("expected data, got %q", got)
	}
}
