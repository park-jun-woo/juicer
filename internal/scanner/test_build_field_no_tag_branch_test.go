//ff:func feature=scan type=test control=sequence
//ff:what TestBuildField_NoTagBranch 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestBuildField_NoTagBranch(t *testing.T) {
	v := types.NewVar(0, nil, "Age", types.Typ[types.Int])
	f := buildField(v, "", make(map[string]bool))
	if f == nil {
		t.Fatal("expected non-nil")
	}
}
