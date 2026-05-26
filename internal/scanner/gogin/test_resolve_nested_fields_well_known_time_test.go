//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_WellKnownTime time.Time Named 타입이 struct 전개되지 않고 nil을 반환하는지 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_WellKnownTime(t *testing.T) {
	timePkg := types.NewPackage("time", "time")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, timePkg, "wall", types.Typ[types.Uint64]),
		types.NewVar(0, timePkg, "ext", types.Typ[types.Int64]),
	}, nil)
	tn := types.NewTypeName(0, timePkg, "Time", nil)
	named := types.NewNamed(tn, st, nil)

	visited := make(map[string]bool)
	fields := resolveNestedFields(named, visited)
	if fields != nil {
		t.Fatalf("expected nil for time.Time (well-known type), got %d fields", len(fields))
	}
}
