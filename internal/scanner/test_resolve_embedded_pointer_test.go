//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveEmbedded_Pointer 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_Pointer(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Base", nil), st, nil)
	ptr := types.NewPointer(named)

	visited := make(map[string]bool)
	fields := resolveEmbedded(ptr, visited)
	if len(fields) != 1 {
		t.Errorf("expected 1 field, got %d", len(fields))
	}
}
