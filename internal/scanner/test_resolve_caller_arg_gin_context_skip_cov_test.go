//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_GinContextSkipCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_GinContextSkipCov(t *testing.T) {
	pkg := types.NewPackage("github.com/gin-gonic/gin", "gin")
	tn := types.NewTypeName(0, pkg, "Context", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	ty := types.NewPointer(named)
	r := resolveCallerArg(ty, &ast.Ident{Name: "c"}, nil)
	if !r.skip {
		t.Fatal("expected skip for *gin.Context")
	}
}
