//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_SliceType 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_SliceType(t *testing.T) {
	result := resolveNestedFields(types.NewSlice(types.Typ[types.String]), make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0 for non-struct slice, got %d", len(result))
	}
}
