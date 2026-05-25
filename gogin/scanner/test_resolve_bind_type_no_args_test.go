//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveBindType_NoArgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveBindType_NoArgs(t *testing.T) {
	call := &ast.CallExpr{}
	typeName, fields := resolveBindType(call, nil)
	if typeName != "" || fields != nil {
		t.Error("expected empty for no args")
	}
}
