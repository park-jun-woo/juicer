//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_PointerType 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_PointerType(t *testing.T) {
	result := resolveNestedFields(types.NewPointer(types.Typ[types.Int]), make(map[string]bool))
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
