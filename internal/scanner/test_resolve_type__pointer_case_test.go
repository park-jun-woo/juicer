//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveType_PointerCase 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_PointerCase(t *testing.T) {
	tn, _ := resolveType(types.NewPointer(types.Typ[types.Int]))
	if tn != "" {
		t.Fatalf("expected empty, got %s", tn)
	}
}
