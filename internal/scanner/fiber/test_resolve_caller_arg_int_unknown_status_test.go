//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_IntUnknownStatus 테스트
package fiber

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_IntUnknownStatus(t *testing.T) {
	intType := types.Typ[types.Int]

	res := resolveCallerArg(intType, &ast.Ident{Name: "statusVar"}, newEmptyInfo())
	if res.status != "" {
		t.Fatalf("expected empty status, got %+v", res)
	}
}
