//ff:func feature=scan type=test control=sequence topic=hono
//ff:what parseBodyResponse / parseJsonResponse 테스트
package hono

import "testing"

func TestParseBodyResponse(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null, 204);`)
	r := parseBodyResponse(call, src)
	if r == nil || r.Kind != "empty" || r.Status != "204" {
		t.Fatalf("got %+v", r)
	}
}

func TestParseBodyResponse_Default(t *testing.T) {
	call, src := firstCallExpr(t, `c.body(null);`)
	r := parseBodyResponse(call, src)
	if r == nil || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}

func TestParseJsonResponse(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({}, 201);`)
	r := parseJsonResponse(call, src)
	if r == nil || r.Kind != "json" || r.Status != "201" {
		t.Fatalf("got %+v", r)
	}
}

func TestParseJsonResponse_Default(t *testing.T) {
	call, src := firstCallExpr(t, `c.json({});`)
	r := parseJsonResponse(call, src)
	if r == nil || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}
