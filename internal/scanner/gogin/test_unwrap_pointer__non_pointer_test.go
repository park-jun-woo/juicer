//ff:func feature=scan type=extract control=sequence
//ff:what TestUnwrapPointer_NonPointer 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestUnwrapPointer_NonPointer(t *testing.T) {
	ty := types.Typ[types.String]
	got := unwrapPointer(ty)
	if got != ty {
		t.Fatal("expected same type")
	}
}
