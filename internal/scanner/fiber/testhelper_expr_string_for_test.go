//ff:func feature=scan type=test control=sequence
//ff:what exprStringFor 테스트 헬퍼
package fiber

import (
	"go/parser"
	"testing"
)

func exprStringFor(t *testing.T, expr string) string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return exprString(e)
}
