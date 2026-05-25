//ff:func feature=scan type=extract control=sequence
//ff:what TestBindVarName_UnaryAnd 테스트
package scanner

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
}
