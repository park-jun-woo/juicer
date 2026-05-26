//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextTypeInfo_CorrectCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_CorrectCov(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	tn := types.NewTypeName(0, pkg, "Context", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if !isGinContextTypeInfo(types.NewPointer(named)) {
		t.Fatal("expected true for *gin.Context")
	}
}
