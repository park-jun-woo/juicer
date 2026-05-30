//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_NotStruct 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_NotStruct(t *testing.T) {

	if got := resolveEmbedded(types.Typ[types.Int], map[string]bool{}); got != nil {
		t.Fatalf("expected nil for non-struct, got %v", got)
	}
}
