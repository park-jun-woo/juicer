//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextTypeInfo_PtrToBasicCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_PtrToBasicCov(t *testing.T) {
	if isGinContextTypeInfo(types.NewPointer(types.Typ[types.Int])) {
		t.Fatal("expected false for ptr-to-basic")
	}
}
