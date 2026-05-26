//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_NonStruct 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_NonStruct(t *testing.T) {
	typeName, fields := resolveType(types.Typ[types.Int])
	if typeName != "" {
		t.Errorf("expected empty typeName, got %q", typeName)
	}
	if fields != nil {
		t.Error("expected nil fields")
	}
}
