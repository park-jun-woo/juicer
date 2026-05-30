//ff:func feature=scan type=test control=sequence
//ff:what TestResolveNestedFields_NonStruct 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_NonStruct(t *testing.T) {
	if got := resolveNestedFields(types.Typ[types.Int], map[string]bool{}); got != nil {
		t.Fatalf("int should yield nil, got %v", got)
	}
}
