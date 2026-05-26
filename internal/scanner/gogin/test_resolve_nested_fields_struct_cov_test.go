//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_StructCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_StructCov(t *testing.T) {
	fields := []*types.Var{types.NewVar(0, nil, "ID", types.Typ[types.Int])}
	st := types.NewStruct(fields, []string{`json:"id"`})
	result := resolveNestedFields(st, make(map[string]bool))
	if len(result) < 1 {
		t.Fatal("expected fields")
	}
}
