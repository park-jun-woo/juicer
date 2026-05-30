//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveTargetEchoAlias_Empty_Round5 테스트
package echo

import (
	"go/token"
	"testing"
)

func TestResolveTargetEchoAlias_Empty_Round5(t *testing.T) {
	if got := resolveTargetEchoAlias(token.Pos(1), emptyGroupCtx()); got != "" {
		t.Fatalf("expected empty alias, got %q", got)
	}
}
