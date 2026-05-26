//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_PointerCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_PointerCov(t *testing.T) {
	resolveType(types.NewPointer(types.Typ[types.Int]))
}
