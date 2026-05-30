//ff:func feature=scan type=test control=sequence topic=echo
//ff:what callWithStringArg 테스트 헬퍼
package echo

import (
	"go/ast"
	"testing"
)

// callWithStringArg builds an *ast.CallExpr whose first arg is the given literal.
func callWithStringArg(t *testing.T, lit string) *ast.CallExpr {
	t.Helper()
	expr := parseExpr(t, "f("+lit+")")
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		t.Fatalf("not a call expr: %s", lit)
	}
	return call
}
