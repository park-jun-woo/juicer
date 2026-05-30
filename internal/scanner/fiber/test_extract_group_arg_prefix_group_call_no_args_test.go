//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_GroupCallNoArgs 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractGroupArgPrefix_GroupCallNoArgs(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Group()`)
	prefix, _, ok := extractGroupArgPrefix(e, groupCtx())
	if !ok || prefix != "/api" {
		t.Fatalf("group call no args: prefix=%q ok=%v", prefix, ok)
	}
}
