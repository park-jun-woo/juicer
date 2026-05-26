//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_NilPkg 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_NilPkg(t *testing.T) {
	named := types.NewNamed(types.NewTypeName(0, nil, "RouterGroup", nil), types.Typ[types.Int], nil)
	pt := types.NewPointer(named)
	if isGinRouterTypeInfo(pt) {
		t.Fatal("expected false for nil package")
	}
}
