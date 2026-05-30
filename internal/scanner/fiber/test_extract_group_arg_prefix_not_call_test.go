//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_NotCall 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractGroupArgPrefix_NotCall(t *testing.T) {

	e, _ := parser.ParseExpr(`42`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("basic lit should be false")
	}
}
