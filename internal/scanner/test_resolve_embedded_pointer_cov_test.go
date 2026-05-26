//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_PointerCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_PointerCov(t *testing.T) {
	fields := []*types.Var{types.NewVar(0, nil, "X", types.Typ[types.Int])}
	st := types.NewStruct(fields, []string{""})
	result := resolveEmbedded(types.NewPointer(st), make(map[string]bool))
	if len(result) < 1 {
		t.Fatal("expected fields from pointer deref")
	}
}
