//ff:func feature=scan type=test control=sequence
//ff:what TestUnwrapPointer_NonPointerCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer_NonPointerCov(t *testing.T) {
	got := unwrapPointer(types.Typ[types.String])
	if got != types.Typ[types.String] {
		t.Fatal("expected same type")
	}
}
