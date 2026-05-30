//ff:func feature=scan type=test control=sequence
//ff:what extractPathFor 테스트 헬퍼
package fiber

import (
	"go/parser"
	"testing"
)

func extractPathFor(t *testing.T, expr string) (string, bool) {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return extractPathString(e)
}
