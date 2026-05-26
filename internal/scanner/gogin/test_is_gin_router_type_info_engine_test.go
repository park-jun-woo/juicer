//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_PointerToGinEngine 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_PointerToGinEngine(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	named := types.NewNamed(types.NewTypeName(0, pkg, "Engine", nil), types.Typ[types.Int], nil)
	pt := types.NewPointer(named)
	if !isGinRouterTypeInfo(pt) {
		t.Fatal("expected true for *gin.Engine")
	}
}
