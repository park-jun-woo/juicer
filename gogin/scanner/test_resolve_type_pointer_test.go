//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_Pointer 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_Pointer(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "Name", types.Typ[types.String]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "User", nil), st, nil)
	ptr := types.NewPointer(named)

	typeName, fields := resolveType(ptr)
	if typeName != "User" {
		t.Errorf("expected 'User', got %q", typeName)
	}
	if len(fields) != 1 {
		t.Errorf("expected 1 field, got %d", len(fields))
	}
}
