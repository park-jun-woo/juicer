//ff:func feature=scan type=test control=sequence
//ff:what parseCall 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func parseCall(t *testing.T, expr string) *ast.CallExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	call, ok := e.(*ast.CallExpr)
	if !ok {
		t.Fatalf("%q is not a call", expr)
	}
	return call
}
