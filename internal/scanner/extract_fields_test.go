//ff:func feature=scan type=test control=sequence
//ff:what TestExtractFields_Basic 테스트
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

