//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveNestedFields_CycleDetection 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveNestedFields_CycleDetection(t *testing.T) {
	pkg := types.NewPackage("test", "test")
	st := types.NewStruct([]*types.Var{
		types.NewVar(0, pkg, "ID", types.Typ[types.Int]),
	}, []string{""})
	named := types.NewNamed(types.NewTypeName(0, pkg, "Node", nil), st, nil)

	visited := map[string]bool{named.String(): true}
	fields := resolveNestedFields(named, visited)
	if fields != nil {
		t.Error("expected nil for cycle")
	}
}
