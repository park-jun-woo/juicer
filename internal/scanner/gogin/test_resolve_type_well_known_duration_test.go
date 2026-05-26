//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_WellKnownDuration time.Duration이 struct 전개 없이 반환되는지 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_WellKnownDuration(t *testing.T) {
	timePkg := types.NewPackage("time", "time")
	tn := types.NewTypeName(0, timePkg, "Duration", nil)
	named := types.NewNamed(tn, types.Typ[types.Int64], nil)

	typeName, gotFields := resolveType(named)
	if typeName != "time.Duration" {
		t.Fatalf("expected time.Duration, got %s", typeName)
	}
	if gotFields != nil {
		t.Fatal("expected nil fields for well-known type")
	}
}
