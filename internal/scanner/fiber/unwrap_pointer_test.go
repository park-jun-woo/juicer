//ff:func feature=scan type=test control=sequence
//ff:what unwrapPointer — 포인터 언래핑 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer(t *testing.T) {
	intT := types.Typ[types.Int]
	ptr := types.NewPointer(intT)
	if got := unwrapPointer(ptr); got != intT {
		t.Errorf("pointer unwrap = %v, want int", got)
	}
	// non-pointer returned as-is
	if got := unwrapPointer(intT); got != intT {
		t.Errorf("non-pointer = %v", got)
	}
}
