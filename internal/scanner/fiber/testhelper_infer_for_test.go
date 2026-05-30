//ff:func feature=scan type=test control=sequence
//ff:what inferFor 테스트 헬퍼
package fiber

import (
	"go/parser"
	"testing"
)

func inferFor(t *testing.T, expr string) string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return inferValueType(e, nil)
}
