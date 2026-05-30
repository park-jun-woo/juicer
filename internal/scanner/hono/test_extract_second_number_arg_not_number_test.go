//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractSecondNumberArg_NotNumber 테스트
package hono

import "testing"

func TestExtractSecondNumberArg_NotNumber(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, opts);`)
	if got := extractSecondNumberArg(call, src, "200"); got != "200" {
		t.Fatalf("got %q", got)
	}
}
