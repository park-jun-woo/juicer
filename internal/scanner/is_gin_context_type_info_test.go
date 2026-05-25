package scanner

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_NonPointer(t *testing.T) {
	if isGinContextTypeInfo(types.Typ[types.String]) {
		t.Fatal("expected false for non-pointer")
	}
}

func TestIsGinContextTypeInfo_PointerToBasic(t *testing.T) {
	if isGinContextTypeInfo(types.NewPointer(types.Typ[types.Int])) {
		t.Fatal("expected false for pointer to basic")
	}
}
