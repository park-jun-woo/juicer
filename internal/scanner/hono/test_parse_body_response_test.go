//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseBodyResponse 테스트
package hono

import "testing"

func TestParseBodyResponse(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null, 204);`)
	r := parseBodyResponse(call, src)
	if r == nil || r.Kind != "empty" || r.Status != "204" {
		t.Fatalf("got %+v", r)
	}
}
