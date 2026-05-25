package scanner

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_NonStructType(t *testing.T) {
	result := resolveEmbedded(types.Typ[types.String], make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}

func TestResolveEmbedded_PointerType(t *testing.T) {
	result := resolveEmbedded(types.NewPointer(types.Typ[types.Int]), make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
