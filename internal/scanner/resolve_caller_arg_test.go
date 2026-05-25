//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveCallerArg_GinContext 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_GinContext(t *testing.T) {
	// Non-gin pointer type should not be gin context
	ty := types.NewPointer(types.Typ[types.String])
	r := resolveCallerArg(ty, &ast.Ident{Name: "c"}, nil)
	if r.skip {
		t.Fatal("should not skip for non-gin pointer")
	}
}
