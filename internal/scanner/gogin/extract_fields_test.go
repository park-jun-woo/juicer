//ff:func feature=scan type=test control=sequence
//ff:what TestExtractFields_Basic 테스트
package gogin

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

	// with embedded field
	embedded := types.NewStruct([]*types.Var{
		types.NewVar(0, nil, "ID", types.Typ[types.Int]),
	}, []string{`json:"id"`})
	namedEmbedded := types.NewNamed(types.NewTypeName(0, nil, "Base", nil), embedded, nil)
	fields2 := []*types.Var{
		types.NewField(0, nil, "Base", namedEmbedded, true),
		types.NewVar(0, nil, "Name", types.Typ[types.String]),
	}
	st2 := types.NewStruct(fields2, []string{"", `json:"name"`})
	result2 := extractFields(st2, make(map[string]bool))
	if len(result2) < 1 {
		t.Fatal("expected at least 1 field with embedded")
	}
}

