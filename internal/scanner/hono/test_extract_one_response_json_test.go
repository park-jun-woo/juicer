//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneResponse_Json 테스트
package hono

import "testing"

func TestExtractOneResponse_Json(t *testing.T) {
	fi := firstCall(t, `c.json({ ok: true });`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "json" || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}
