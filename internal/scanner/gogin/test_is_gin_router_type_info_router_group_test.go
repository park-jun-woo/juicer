//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_PointerToGinRouterGroup 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_PointerToGinRouterGroup(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	named := types.NewNamed(types.NewTypeName(0, pkg, "RouterGroup", nil), types.Typ[types.Int], nil)
	pt := types.NewPointer(named)
	if !isGinRouterTypeInfo(pt) {
		t.Fatal("expected true for *gin.RouterGroup")
	}
}
