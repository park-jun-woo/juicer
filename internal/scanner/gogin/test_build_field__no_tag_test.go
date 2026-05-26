//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildField_NoTag 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestBuildField_NoTag(t *testing.T) {
	v := types.NewVar(0, nil, "ID", types.Typ[types.Int])
	f := buildField(v, "", make(map[string]bool))
	if f == nil {
		t.Fatal("expected non-nil")
	}
}
