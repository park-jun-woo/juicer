//ff:func feature=scan type=test control=sequence
//ff:what collectFor 테스트 헬퍼
package fiber

import (
	"go/parser"
	"testing"
)

func collectFor(t *testing.T, expr string) []string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	var parts []string
	collectStringParts(e, &parts)
	return parts
}
