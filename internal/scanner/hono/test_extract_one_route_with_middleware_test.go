//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_WithMiddleware 테스트
package hono

import "testing"

func TestExtractOneRoute_WithMiddleware(t *testing.T) {
	r := oneRoute(t, `app.post("/x", auth, handler);`, map[string]bool{"app": true})
	if r == nil || len(r.Middleware) != 1 || r.Middleware[0] != "auth" {
		t.Fatalf("got %+v", r)
	}
	if r.Line != 1 {
		t.Fatalf("expected line 1, got %d", r.Line)
	}
}
