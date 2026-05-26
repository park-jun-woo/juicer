//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_Slice 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_Slice(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Item", nil), st, nil)
	sl := types.NewSlice(named)

	typeName, fields := resolveType(sl)
	if typeName != "[]Item" {
		t.Errorf("expected '[]Item', got %q", typeName)
	}
	if len(fields) != 1 {
		t.Errorf("expected 1 field, got %d", len(fields))
	}
}
