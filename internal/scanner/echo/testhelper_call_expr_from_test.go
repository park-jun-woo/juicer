//ff:func feature=scan type=test control=sequence topic=echo
//ff:what callExprFrom 테스트 헬퍼
package echo

import (
	"go/ast"
	"testing"
)

// callExprFrom parses an expression and returns it as *ast.CallExpr.
func callExprFrom(t *testing.T, src string) *ast.CallExpr {
	t.Helper()
	expr := parseExpr(t, src)
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		t.Fatalf("not a call: %s", src)
	}
	return call
}
