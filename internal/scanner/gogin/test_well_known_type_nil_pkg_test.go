//ff:func feature=scan type=test control=sequence
//ff:what TestWellKnownType_NilPkg 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestWellKnownType_NilPkg(t *testing.T) {

	tn := types.NewTypeName(0, nil, "Ctx", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if _, ok := wellKnownType(named); ok {
		t.Fatal("nil-pkg type should not be well-known")
	}
}
