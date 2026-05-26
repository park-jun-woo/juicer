//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextTypeInfo_NilPkgCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestIsGinContextTypeInfo_NilPkgCov(t *testing.T) {
	tn := types.NewTypeName(0, nil, "Context", nil)
	named := types.NewNamed(tn, types.Typ[types.Int], nil)
	if isGinContextTypeInfo(types.NewPointer(named)) {
		t.Fatal("expected false for nil pkg")
	}
}
