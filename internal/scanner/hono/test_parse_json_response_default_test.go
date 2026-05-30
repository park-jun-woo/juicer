//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseJsonResponse_Default 테스트
package hono

import "testing"

func TestParseJsonResponse_Default(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({});`)
	r := parseJsonResponse(call, src)
	if r == nil || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}
