//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_NotWellKnown 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestWellKnownType_NotWellKnown(t *testing.T) {
	pkg := types.NewPackage("mypkg", "mypkg")
	tn := types.NewTypeName(0, pkg, "Custom", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if _, ok := wellKnownType(named); ok {
		t.Fatal("custom type should not be well-known")
	}
}
