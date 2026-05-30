//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneResponse_JsonStatus 테스트
package hono

import "testing"

func TestExtractOneResponse_JsonStatus(t *testing.T) {
	fi := firstCall(t, `c.json({ ok: true }, 201);`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Status != "201" {
		t.Fatalf("got %+v", r)
	}
}
