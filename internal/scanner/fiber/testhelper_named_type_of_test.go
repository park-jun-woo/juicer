//ff:func feature=scan type=test control=sequence
//ff:what namedTypeOf 테스트 헬퍼
package fiber

import (
	"go/types"
	"testing"
)

func namedTypeOf(t *testing.T, src, varName string) *types.Named {
	t.Helper()
	typ := typeOfVar(t, src, varName)
	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}
	named, ok := typ.(*types.Named)
	if !ok {
		t.Fatalf("%s is not a named type: %T", varName, typ)
	}
	return named
}
