//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_WellKnownTime time.Time이 struct 전개 없이 "time.Time"으로 반환되는지 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_WellKnownTime(t *testing.T) {
	timePkg := types.NewPackage("time", "time")
	fields := []*types.Var{
		types.NewVar(0, timePkg, "wall", types.Typ[types.Uint64]),
		types.NewVar(0, timePkg, "ext", types.Typ[types.Int64]),
	}
	st := types.NewStruct(fields, nil)
	tn := types.NewTypeName(0, timePkg, "Time", nil)
	named := types.NewNamed(tn, st, nil)

	typeName, gotFields := resolveType(named)
	if typeName != "time.Time" {
		t.Fatalf("expected time.Time, got %s", typeName)
	}
	if gotFields != nil {
		t.Fatal("expected nil fields for well-known type, got struct expansion")
	}
}
