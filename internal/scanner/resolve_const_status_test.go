package scanner

import (
	"go/constant"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveConstStatus_Const(t *testing.T) {
	c := types.NewConst(token.NoPos, nil, "StatusOK", types.Typ[types.Int], constant.MakeInt64(200))
	got := resolveConstStatus(c)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}

func TestResolveConstStatus_NonConst(t *testing.T) {
	v := types.NewVar(token.NoPos, nil, "x", types.Typ[types.Int])
	got := resolveConstStatus(v)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
