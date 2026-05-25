//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_Array 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_Array(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Item", nil), st, nil)
	arr := types.NewArray(named, 5)

	typeName, fields := resolveType(arr)
	if typeName != "[]Item" {
		t.Errorf("expected '[]Item', got %q", typeName)
	}
	if len(fields) != 1 {
		t.Errorf("expected 1 field, got %d", len(fields))
	}
}
