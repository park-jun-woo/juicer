//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_Basic 테스트
package gogin

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

	// json:"-" should return nil
	v2 := types.NewVar(0, nil, "Secret", types.Typ[types.String])
	f2 := buildField(v2, `json:"-"`, make(map[string]bool))
	if f2 != nil {
		t.Fatal("expected nil for json:\"-\"")
	}
}
