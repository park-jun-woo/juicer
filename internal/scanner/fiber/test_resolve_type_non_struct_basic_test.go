//ff:func feature=scan type=test control=sequence
//ff:what TestResolveType_NonStructBasic 테스트
package fiber

import (
	"go/types"
	"testing"
)

func TestResolveType_NonStructBasic(t *testing.T) {
	tn, f := resolveType(types.Typ[types.String])
	if tn != "" || f != nil {
		t.Fatalf("basic string: %q %v", tn, f)
	}
}
