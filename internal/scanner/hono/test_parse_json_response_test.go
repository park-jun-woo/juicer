//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestParseJsonResponse 테스트
package hono

import "testing"

func TestParseJsonResponse(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, 201);`)
	r := parseJsonResponse(call, src)
	if r == nil || r.Kind != "json" || r.Status != "201" {
		t.Fatalf("got %+v", r)
	}
}
