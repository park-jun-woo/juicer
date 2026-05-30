//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseBodyResponse_Default 테스트
package hono

import "testing"

func TestParseBodyResponse_Default(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null);`)
	r := parseBodyResponse(call, src)
	if r == nil || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}
