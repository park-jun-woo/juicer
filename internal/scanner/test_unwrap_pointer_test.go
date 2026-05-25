//ff:func feature=scan type=extract control=sequence
//ff:what TestUnwrapPointer 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer(t *testing.T) {
	basic := types.Typ[types.Int]
	ptr := types.NewPointer(basic)

	if unwrapPointer(ptr) != basic {
		t.Error("expected unwrapped pointer to be basic int")
	}
	if unwrapPointer(basic) != basic {
		t.Error("expected non-pointer to be returned as-is")
	}
}
