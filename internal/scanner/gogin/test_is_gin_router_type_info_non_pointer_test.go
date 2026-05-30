//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_NonPointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_NonPointer(t *testing.T) {
	basic := types.Typ[types.Int]
	if isGinRouterTypeInfo(basic) {
		t.Fatal("expected false for non-pointer type")
	}
}
