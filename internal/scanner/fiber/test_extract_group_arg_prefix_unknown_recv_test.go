//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_UnknownRecv 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractGroupArgPrefix_UnknownRecv(t *testing.T) {
	e, _ := parser.ParseExpr(`unknown.Group("/x")`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("unknown receiver should be false")
	}
}
