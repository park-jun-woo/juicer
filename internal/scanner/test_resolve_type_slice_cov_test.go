//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_SliceCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_SliceCov(t *testing.T) {
	f := []*types.Var{types.NewVar(0, nil, "X", types.Typ[types.Int])}
	st := types.NewStruct(f, []string{""})
	tn := types.NewTypeName(0, nil, "Item", nil)
	named := types.NewNamed(tn, st, nil)
	typeName, _ := resolveType(types.NewSlice(named))
	if typeName != "[]Item" {
		t.Fatalf("expected []Item, got %s", typeName)
	}
}
