//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextTypeInfo_WrongNameCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_WrongNameCov(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	tn := types.NewTypeName(0, pkg, "Engine", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if isGinContextTypeInfo(types.NewPointer(named)) {
		t.Fatal("expected false for wrong name")
	}
}
