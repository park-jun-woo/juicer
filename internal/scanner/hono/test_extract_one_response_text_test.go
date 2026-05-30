//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneResponse_Text 테스트
package hono

import "testing"

func TestExtractOneResponse_Text(t *testing.T) {
	fi := firstCall(t, `c.text("hi");`)
	call := findAllByType(fi.Root, "call_expression")[0]
	r := extractOneResponse(call, fi.Src)
	if r == nil || r.Kind != "text" || r.Status != "200" {
		t.Fatalf("got %+v", r)
	}
}
