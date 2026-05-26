//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_SliceCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_SliceCov(t *testing.T) {
	fields := []*types.Var{types.NewVar(0, nil, "X", types.Typ[types.Int])}
	st := types.NewStruct(fields, []string{""})
	result := resolveNestedFields(types.NewSlice(st), make(map[string]bool))
	if len(result) < 1 {
		t.Fatal("expected fields from slice elem")
	}
}
