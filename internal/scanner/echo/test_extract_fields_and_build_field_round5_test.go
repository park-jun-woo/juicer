//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestExtractFields_And_BuildField_Round5 테스트
package echo

import (
	"go/types"
	"testing"
)

func TestExtractFields_And_BuildField_Round5(t *testing.T) {
	_, info := checkSrc(t, dtoSrc)
	typ := namedType(t, info, "U")
	named := typ.(*types.Named)
	st := named.Underlying().(*types.Struct)
	fields := extractFields(st, map[string]bool{})
	// embedded Base.ID should be flattened in
	var hasID, hasName bool
	for _, f := range fields {
		if f.JSON == "id" {
			hasID = true
		}
		if f.JSON == "name" {
			hasName = true
		}
	}
	if !hasID {
		t.Errorf("embedded id not flattened: %+v", fields)
	}
	if !hasName {
		t.Errorf("name missing: %+v", fields)
	}

	// buildField with json:"-" returns nil
	var hiddenVar *types.Var
	for i := 0; i < st.NumFields(); i++ {
		if st.Field(i).Name() == "Hidden" {
			hiddenVar = st.Field(i)
		}
	}
	if hiddenVar == nil {
		t.Fatal("no hidden field")
	}
	if buildField(hiddenVar, `json:"-"`, map[string]bool{}) != nil {
		t.Error("json:\"-\" field should be nil")
	}
}
