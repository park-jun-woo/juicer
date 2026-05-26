//ff:func feature=scan type=test control=sequence
//ff:what TestExtractFields_EmbeddedCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestExtractFields_EmbeddedCov(t *testing.T) {
	innerFields := []*types.Var{
		types.NewVar(0, nil, "ID", types.Typ[types.Int]),
	}
	innerSt := types.NewStruct(innerFields, []string{`json:"id"`})
	innerNamed := types.NewNamed(types.NewTypeName(0, nil, "Base", nil), innerSt, nil)

	outerFields := []*types.Var{
		types.NewField(0, nil, "Base", innerNamed, true),
		types.NewVar(0, nil, "Name", types.Typ[types.String]),
	}
	outerSt := types.NewStruct(outerFields, []string{"", `json:"name"`})
	result := extractFields(outerSt, make(map[string]bool))
	if len(result) < 1 {
		t.Fatalf("expected at least 1 field, got %d", len(result))
	}
}
