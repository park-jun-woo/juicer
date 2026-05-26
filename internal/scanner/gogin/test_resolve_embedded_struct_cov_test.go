//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_StructCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_StructCov(t *testing.T) {
	fields := []*types.Var{types.NewVar(0, nil, "ID", types.Typ[types.Int])}
	st := types.NewStruct(fields, []string{`json:"id"`})
	tn := types.NewTypeName(0, nil, "Base", nil)
	named := types.NewNamed(tn, st, nil)
	result := resolveEmbedded(named, make(map[string]bool))
	if len(result) < 1 {
		t.Fatalf("expected at least 1 field, got %d", len(result))
	}
}
