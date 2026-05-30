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

func TestWellKnownType_NilPkg(t *testing.T) {
	// type name with nil package (universe scope) -> false
	tn := types.NewTypeName(0, nil, "Ctx", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if _, ok := wellKnownType(named); ok {
		t.Fatal("nil-pkg type should not be well-known")
	}
}

func TestWellKnownType_NotWellKnown(t *testing.T) {
	pkg := types.NewPackage("mypkg", "mypkg")
	tn := types.NewTypeName(0, pkg, "Custom", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if _, ok := wellKnownType(named); ok {
		t.Fatal("custom type should not be well-known")
	}
}
