//ff:func feature=scan type=test control=sequence
//ff:what exprNameFor 테스트 헬퍼
package fiber

import (
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
