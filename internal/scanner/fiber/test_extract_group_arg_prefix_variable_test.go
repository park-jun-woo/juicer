//ff:func feature=scan type=test control=sequence
//ff:what TestExtractGroupArgPrefix_Variable 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestExtractGroupArgPrefix_Variable(t *testing.T) {
	arg := ast.NewIdent("authGroup")
	prefix, ri, ok := extractGroupArgPrefix(arg, groupCtx())
	if !ok || prefix != "/auth" || ri == nil {
		t.Fatalf("variable group: prefix=%q ok=%v", prefix, ok)
	}
}
