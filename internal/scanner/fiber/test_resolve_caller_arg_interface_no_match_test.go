//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_InterfaceNoMatch 테스트
package fiber

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_InterfaceNoMatch(t *testing.T) {

	emptyIface := types.NewInterfaceType(nil, nil)
	res := resolveCallerArg(emptyIface, &ast.Ident{Name: "x"}, newEmptyInfo())
	if res.skip || res.typeName != "" {
		t.Fatalf("expected empty non-skip result, got %+v", res)
	}
}
