//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_ExcludedCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestBuildField_ExcludedCov(t *testing.T) {
	v := types.NewVar(0, nil, "Hidden", types.Typ[types.String])
	f := buildField(v, `json:"-"`, make(map[string]bool))
	if f != nil {
		t.Fatal("expected nil for excluded field")
	}
}
