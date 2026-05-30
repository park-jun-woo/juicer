//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractMiddlewareName_Call 테스트
package hono

import "testing"

func TestExtractMiddlewareName_Call(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", auth(), h);`)
	if got != "auth" {
		t.Fatalf("got %q", got)
	}
}
