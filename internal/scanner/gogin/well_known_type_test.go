//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_TimeTime well-known 타입 판별 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestWellKnownType_TimeTime(t *testing.T) {
	timePkg := types.NewPackage("time", "time")
	tn := types.NewTypeName(0, timePkg, "Time", nil)
	named := types.NewNamed(tn, types.Typ[types.Int64], nil)

	got, ok := wellKnownType(named)
	if !ok || got != "time.Time" {
		t.Fatalf("expected (time.Time, true), got (%s, %v)", got, ok)
	}
}
