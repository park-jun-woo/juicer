//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveEmbedded_NonStructType 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_NonStructType(t *testing.T) {
	result := resolveEmbedded(types.Typ[types.String], make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
