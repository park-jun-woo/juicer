//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_Array 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_Array(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	innerSt := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	innerNamed := types.NewNamed(types.NewTypeName(0, pkg, "Item", nil), innerSt, nil)
	arr := types.NewArray(innerNamed, 3)

	visited := make(map[string]bool)
	fields := resolveNestedFields(arr, visited)
	if len(fields) != 1 {
		t.Errorf("expected 1 nested field, got %d", len(fields))
	}
}
