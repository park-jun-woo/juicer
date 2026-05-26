//ff:func feature=scan type=test control=sequence
//ff:what TestUnwrapPointer_Pointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer_Pointer(t *testing.T) {
	inner := types.Typ[types.Int]
	ptr := types.NewPointer(inner)
	got := unwrapPointer(ptr)
	if got != inner {
		t.Fatal("expected unwrapped type")
	}
}

