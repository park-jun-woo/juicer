//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_SliceOfPointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_SliceOfPointer(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Item", nil), st, nil)
	sl := types.NewSlice(types.NewPointer(named))

	typeName, _ := resolveType(sl)
	if typeName != "[]Item" {
		t.Errorf("expected '[]Item', got %q", typeName)
	}
}
