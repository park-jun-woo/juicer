//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinContextTypeInfo_PointerToBasic 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_PointerToBasic(t *testing.T) {
	if isGinContextTypeInfo(types.NewPointer(types.Typ[types.Int])) {
		t.Fatal("expected false for pointer to basic")
	}
}
