//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneResponse_Body 테스트
package hono

import "testing"

func TestExtractOneResponse_Body(t *testing.T) {
	fi := firstCall(t, `c.body(null, 204);`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "empty" || r.Status != "204" {
		t.Fatalf("got %+v", r)
	}
}
