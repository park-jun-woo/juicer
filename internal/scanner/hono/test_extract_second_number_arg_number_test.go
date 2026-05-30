//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractSecondNumberArg_Number 테스트
package hono

import "testing"

func TestExtractSecondNumberArg_Number(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, 201);`)
	if got := extractSecondNumberArg(call, src, "200"); got != "201" {
		t.Fatalf("got %q", got)
	}
}
