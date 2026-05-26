//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_PointerToWrongPkg 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_PointerToWrongPkg(t *testing.T) {
	pkg := types.NewPackage("example.com/other", "other")
	named := types.NewNamed(types.NewTypeName(0, pkg, "RouterGroup", nil), types.Typ[types.Int], nil)
	pt := types.NewPointer(named)
	if isGinRouterTypeInfo(pt) {
		t.Fatal("expected false for wrong package")
	}
}
