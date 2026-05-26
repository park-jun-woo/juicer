//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_NamedStructCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_NamedStructCov(t *testing.T) {
	f := []*types.Var{types.NewVar(0, nil, "ID", types.Typ[types.Int])}
	st := types.NewStruct(f, []string{`json:"id"`})
	tn := types.NewTypeName(0, nil, "User", nil)
	named := types.NewNamed(tn, st, nil)
	typeName, fields := resolveType(named)
	if typeName != "User" {
		t.Fatalf("expected User, got %s", typeName)
	}
	if len(fields) < 1 {
		t.Fatal("expected fields")
	}
}
