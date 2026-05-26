//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_BasicType 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_BasicType(t *testing.T) {
	result := resolveNestedFields(types.Typ[types.String], make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}

