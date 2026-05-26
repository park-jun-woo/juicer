//ff:func feature=scan type=test control=sequence
//ff:what TestResolveEmbedded_CyclicCov 테스트
package gogin

import (
	"go/types"
	"testing"
)

func TestResolveEmbedded_CyclicCov(t *testing.T) {
	st := types.NewStruct(nil, nil)
	tn := types.NewTypeName(0, nil, "Cycle", nil)
	named := types.NewNamed(tn, st, nil)
	visited := map[string]bool{named.String(): true}
	result := resolveEmbedded(named, visited)
	if result != nil {
		t.Fatal("expected nil for cyclic")
	}
}
