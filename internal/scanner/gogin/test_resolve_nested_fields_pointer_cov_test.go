//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_PointerCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_PointerCov(t *testing.T) {
	result := resolveNestedFields(types.NewPointer(types.Typ[types.Int]), make(map[string]bool))
	if result != nil {
		t.Fatal("expected nil for pointer-to-basic")
	}
}
