//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveEmbedded_NonStruct 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_NonStruct(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	named := types.NewNamed(types.NewTypeName(0, pkg, "MyInt", nil), types.Typ[types.Int], nil)
	visited := make(map[string]bool)
	fields := resolveEmbedded(named, visited)
	if fields != nil {
		t.Error("expected nil for non-struct embedded type")
	}
}
