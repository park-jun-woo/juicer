//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_NotGroupSelector 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractGroupArgPrefix_NotGroupSelector(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Use(mw)`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("non-Group selector should be false")
	}
}
