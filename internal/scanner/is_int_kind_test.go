package scanner

import (
	"go/types"
	"testing"
)

func TestIsIntKind_Int(t *testing.T) {
	if !isIntKind(types.Int) {
		t.Fatal("expected true")
	}
}

func TestIsIntKind_Int64(t *testing.T) {
	if !isIntKind(types.Int64) {
		t.Fatal("expected true")
	}
}

func TestIsIntKind_String(t *testing.T) {
	if isIntKind(types.String) {
		t.Fatal("expected false")
	}
}

func TestIsIntKind_Uint(t *testing.T) {
	if !isIntKind(types.Uint) {
		t.Fatal("expected true")
	}
}
