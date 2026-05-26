//ff:func feature=scan type=test control=sequence
//ff:what TestResolveConstStatus_Const 테스트
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
