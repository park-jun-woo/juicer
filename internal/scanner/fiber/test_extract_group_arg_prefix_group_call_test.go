//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_GroupCall 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractGroupArgPrefix_GroupCall(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Group("/v1")`)
	prefix, ri, ok := extractGroupArgPrefix(e, groupCtx())
	if !ok || prefix != "/api/v1" || ri == nil {
		t.Fatalf("group call: prefix=%q ok=%v", prefix, ok)
	}
}
