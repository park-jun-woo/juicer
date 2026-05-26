//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_WellKnownTimeSlice []time.Time이 struct 전개 없이 반환되는지 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_WellKnownTimeSlice(t *testing.T) {
	timePkg := types.NewPackage("time", "time")
	fields := []*types.Var{
		types.NewVar(0, timePkg, "wall", types.Typ[types.Uint64]),
	}
	st := types.NewStruct(fields, nil)
	tn := types.NewTypeName(0, timePkg, "Time", nil)
	named := types.NewNamed(tn, st, nil)
	sl := types.NewSlice(named)

	typeName, gotFields := resolveType(sl)
	if typeName != "[]time.Time" {
		t.Fatalf("expected []time.Time, got %s", typeName)
	}
	if gotFields != nil {
		t.Fatal("expected nil fields for well-known slice type")
	}
}
