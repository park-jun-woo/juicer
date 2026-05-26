//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_NamedInterfaceWrongPkg 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_NamedInterfaceWrongPkg(t *testing.T) {
	pkg := types.NewPackage("example.com/other", "other")
	iface := types.NewInterfaceType(nil, nil)
	named := types.NewNamed(types.NewTypeName(0, pkg, "IRouter", nil), iface, nil)
	if isGinRouterTypeInfo(named) {
		t.Fatal("expected false for other.IRouter (wrong package)")
	}
}
