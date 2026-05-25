package scanner

import (
	"go/types"
	"testing"
)

func TestBuildField_Basic(t *testing.T) {
	v := types.NewVar(0, nil, "Name", types.Typ[types.String])
	f := buildField(v, `json:"name"`, make(map[string]bool))
	if f == nil {
		t.Fatal("expected non-nil")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}
}

func TestBuildField_Excluded(t *testing.T) {
	v := types.NewVar(0, nil, "Secret", types.Typ[types.String])
	f := buildField(v, `json:"-"`, make(map[string]bool))
	if f != nil {
		t.Fatal("expected nil for json:-")
	}
}

func TestBuildField_NoTag(t *testing.T) {
	v := types.NewVar(0, nil, "ID", types.Typ[types.Int])
	f := buildField(v, "", make(map[string]bool))
	if f == nil {
		t.Fatal("expected non-nil")
	}
}
