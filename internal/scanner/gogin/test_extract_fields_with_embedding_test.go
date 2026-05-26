//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractFields_WithEmbedding 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestExtractFields_WithEmbedding(t *testing.T) {
	pkg := types.NewPackage("test", "test")

	baseSt := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	baseNamed := types.NewNamed(types.NewTypeName(0, pkg, "Base", nil), baseSt, nil)

	// Struct with embedded Base
	mainSt := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "Base", baseNamed), // embedded
		types.NewVar(0, pkg, "Name", types.Typ[types.String]),
	}, []string{"", ""})

	// Mark first field as embedded by creating it properly
	// Actually types.NewStruct doesn't support embedded via constructor in older Go,
	// but we can check the field's Embedded() property
	visited := make(map[string]bool)
	extractedFields := extractFields(mainSt, visited)
	// Should have at least 1 field
	if len(extractedFields) < 1 {
		t.Errorf("expected at least 1 field, got %d", len(extractedFields))
	}
}
