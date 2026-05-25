package scanner

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

func TestUnwrapPointer_NonPointer(t *testing.T) {
	ty := types.Typ[types.String]
	got := unwrapPointer(ty)
	if got != ty {
		t.Fatal("expected same type")
	}
}
