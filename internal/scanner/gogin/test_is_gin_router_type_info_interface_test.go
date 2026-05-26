//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_NamedInterface 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_NamedInterface(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	iface := types.NewInterfaceType(nil, nil)
	named := types.NewNamed(types.NewTypeName(0, pkg, "IRouter", nil), iface, nil)
	if !isGinRouterTypeInfo(named) {
		t.Fatal("expected true for gin.IRouter (named interface)")
	}
}
