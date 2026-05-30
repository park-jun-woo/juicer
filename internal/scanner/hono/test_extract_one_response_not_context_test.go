//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneResponse_NotContext 테스트
package hono

import "testing"

func TestExtractOneResponse_NotContext(t *testing.T) {

	fi := firstCall(t, `res.json({});`)
	call := findAllByType(fi.Root, "call_expression")[0]
	if r := extractOneResponse(call, fi.Src); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
