//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveConstStatus_NonConst 테스트
package gogin

import (
	"go/token"
	"go/types"
	"testing"
)

func TestResolveConstStatus_NonConst(t *testing.T) {
	v := types.NewVar(token.NoPos, nil, "x", types.Typ[types.Int])
	got := resolveConstStatus(v)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
