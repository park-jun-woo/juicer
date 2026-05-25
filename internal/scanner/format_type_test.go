package scanner

import (
	"go/types"
	"testing"
)

func TestFormatType_Basic(t *testing.T) {
	got := formatType(types.Typ[types.String])
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}

func TestFormatType_Pointer(t *testing.T) {
	got := formatType(types.NewPointer(types.Typ[types.Int]))
	if got != "*int" {
		t.Fatalf("expected *int, got %s", got)
	}
}

func TestFormatType_Slice(t *testing.T) {
	got := formatType(types.NewSlice(types.Typ[types.String]))
	if got != "[]string" {
		t.Fatalf("expected []string, got %s", got)
	}
}

func TestFormatType_Interface(t *testing.T) {
	got := formatType(types.NewInterfaceType(nil, nil))
	if got != "any" {
		t.Fatalf("expected any, got %s", got)
	}
}

func TestFormatType_Map(t *testing.T) {
	got := formatType(types.NewMap(types.Typ[types.String], types.Typ[types.Int]))
	if got != "map[string]int" {
		t.Fatalf("expected map[string]int, got %s", got)
	}
}

func TestFormatType_StructType(t *testing.T) {
	got := formatType(types.NewStruct(nil, nil))
	if got != "object" {
		t.Fatalf("expected object, got %s", got)
	}
}

func TestFormatType_Array(t *testing.T) {
	got := formatType(types.NewArray(types.Typ[types.Int], 5))
	if got != "[]int" {
		t.Fatalf("expected []int, got %s", got)
	}
}
