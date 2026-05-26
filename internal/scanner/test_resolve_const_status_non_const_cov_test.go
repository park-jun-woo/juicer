//ff:func feature=scan type=test control=sequence
//ff:what TestResolveConstStatus_NonConstCov 테스트
package scanner

import (
	"go/types"
	"testing"
)

func TestResolveConstStatus_NonConstCov(t *testing.T) {
	v := types.NewVar(0, nil, "x", types.Typ[types.Int])
	got := resolveConstStatus(v)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
