//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_ArrayOfPointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_ArrayOfPointer(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Item", nil), st, nil)
	arr := types.NewArray(types.NewPointer(named), 5)

	typeName, _ := resolveType(arr)
	if typeName != "[]Item" {
		t.Errorf("expected '[]Item', got %q", typeName)
	}
}
