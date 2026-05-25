package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_BasicType(t *testing.T) {
	result := resolveNestedFields(types.Typ[types.String], make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}

func TestResolveNestedFields_PointerType(t *testing.T) {
	result := resolveNestedFields(types.NewPointer(types.Typ[types.Int]), make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}

func TestResolveNestedFields_SliceType(t *testing.T) {
	result := resolveNestedFields(types.NewSlice(types.Typ[types.String]), make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0 for non-struct slice, got %d", len(result))
	}
}
