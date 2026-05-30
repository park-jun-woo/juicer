//ff:func feature=scan type=test control=sequence
//ff:what binExpr 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func binExpr(t *testing.T, expr string) *ast.BinaryExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	be, ok := e.(*ast.BinaryExpr)
	if !ok {
		t.Fatalf("%q is not a binary expr", expr)
	}
	return be
}
