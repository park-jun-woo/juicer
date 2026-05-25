//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_NonStruct 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_NonStruct(t *testing.T) {
	visited := make(map[string]bool)
	fields := resolveNestedFields(types.Typ[types.Int], visited)
	if fields != nil {
		t.Error("expected nil for non-struct")
	}
}
