//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_Basic 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveType_Basic(t *testing.T) {
	tn, fields := resolveType(types.Typ[types.String])
	if tn != "" {
		t.Fatalf("expected empty name for basic type, got %s", tn)
	}
	if fields != nil {
		t.Fatal("expected nil fields for basic type")
	}
}
