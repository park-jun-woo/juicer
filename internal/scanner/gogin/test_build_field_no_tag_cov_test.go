//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_NoTagCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestBuildField_NoTagCov(t *testing.T) {
	v := types.NewVar(0, nil, "Foo", types.Typ[types.Int])
	f := buildField(v, "", make(map[string]bool))
	if f == nil {
		t.Fatal("expected non-nil")
	}
}
