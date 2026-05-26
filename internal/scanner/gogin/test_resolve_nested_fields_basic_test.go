//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_Basic 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_Basic(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	innerSt := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "City", types.Typ[types.String]),
	}, []string{""})
	innerNamed := types.NewNamed(types.NewTypeName(0, pkg, "Address", nil), innerSt, nil)

	visited := make(map[string]bool)
	fields := resolveNestedFields(innerNamed, visited)
	if len(fields) != 1 {
		t.Errorf("expected 1 nested field, got %d", len(fields))
	}
}
