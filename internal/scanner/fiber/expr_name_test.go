//ff:func feature=scan type=test control=selection
//ff:what exprName — 표현 이름 추출 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func exprNameFor(t *testing.T, expr string) string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return exprName(e)
}

func TestExprName(t *testing.T) {
	if got := exprNameFor(t, "handler"); got != "handler" {
		t.Errorf("ident: %q", got)
	}
	if got := exprNameFor(t, "pkg.Handler"); got != "pkg.Handler" {
		t.Errorf("selector with recv: %q", got)
	}
	if got := exprNameFor(t, "func() {}"); got != "(inline)" {
		t.Errorf("funclit: %q", got)
	}
	if got := exprNameFor(t, "make()"); got != "make()" {
		t.Errorf("call: %q", got)
	}
	// default: a non-name expression
	if got := exprNameFor(t, "1 + 2"); got != "" {
		t.Errorf("default: %q, want empty", got)
	}
}

func TestExprName_SelectorNoRecv(t *testing.T) {
	// selector whose X is not a simple ident -> recv "" -> Sel name only
	sel := &ast.SelectorExpr{
		X:   &ast.CallExpr{Fun: ast.NewIdent("f")},
		Sel: ast.NewIdent("Method"),
	}
	if got := exprName(sel); got != "Method" {
		t.Errorf("selector no-recv: %q, want Method", got)
	}
}
