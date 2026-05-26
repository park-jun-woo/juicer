//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_PointerToNonNamed 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_PointerToNonNamed(t *testing.T) {
	pt := types.NewPointer(types.Typ[types.String])
	if isGinRouterTypeInfo(pt) {
		t.Fatal("expected false for pointer to basic type")
	}
}
