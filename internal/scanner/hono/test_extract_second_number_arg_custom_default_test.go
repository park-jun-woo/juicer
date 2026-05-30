//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractSecondNumberArg_CustomDefault 테스트
package hono

import "testing"

func TestExtractSecondNumberArg_CustomDefault(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null);`)
	if got := extractSecondNumberArg(call, src, "204"); got != "204" {
		t.Fatalf("got %q", got)
	}
}
