package scanner

import (
	"go/types"
	"testing"
)

func TestExtractFields_Basic(t *testing.T) {
	fields := []*types.Var{
		types.NewVar(0, nil, "Name", types.Typ[types.String]),
		types.NewVar(0, nil, "Age", types.Typ[types.Int]),
	}
	tags := []string{`json:"name"`, `json:"age"`}
	st := types.NewStruct(fields, tags)
	result := extractFields(st, make(map[string]bool))
	if len(result) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(result))
	}
}

func TestExtractFields_Empty(t *testing.T) {
	st := types.NewStruct(nil, nil)
	result := extractFields(st, make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0 fields, got %d", len(result))
	}
}
