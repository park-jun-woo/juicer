//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_SliceCase 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveType_SliceCase(t *testing.T) {
	tn, _ := resolveType(types.NewSlice(types.Typ[types.String]))
	_ = tn
}
