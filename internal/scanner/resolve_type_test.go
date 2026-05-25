package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_Basic(t *testing.T) {
	tn, fields := resolveType(types.Typ[types.String])
	if tn != "" {
		t.Fatalf("expected empty name for basic type, got %s", tn)
	}
	if fields != nil {
		t.Fatal("expected nil fields for basic type")
	}
}

func TestResolveType_PointerCase(t *testing.T) {
	tn, _ := resolveType(types.NewPointer(types.Typ[types.Int]))
	if tn != "" {
		t.Fatalf("expected empty, got %s", tn)
	}
}

func TestResolveType_SliceCase(t *testing.T) {
	tn, _ := resolveType(types.NewSlice(types.Typ[types.String]))
	_ = tn
}
