//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_ExcludedBranch 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestBuildField_ExcludedBranch(t *testing.T) {
	v := types.NewVar(0, nil, "Internal", types.Typ[types.String])
	f := buildField(v, `json:"-"`, make(map[string]bool))
	if f != nil {
		t.Fatal("expected nil for excluded field")
	}
}
