//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextTypeInfo_NonPointer 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_NonPointer(t *testing.T) {
	if isGinContextTypeInfo(types.Typ[types.String]) {
		t.Fatal("expected false for non-pointer")
	}
}
