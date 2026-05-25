//ff:func feature=scan type=extract control=sequence
//ff:what TestExprName_CallExpr 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprName_CallExpr(t *testing.T) {
	got := exprName(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}})
	if got != "f()" {
		t.Fatalf("expected f(), got %s", got)
	}
}
