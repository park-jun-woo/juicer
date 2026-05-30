//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractMiddlewareName_Other 테스트
package hono

import "testing"

func TestExtractMiddlewareName_Other(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", { a: 1 }, h);`)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
