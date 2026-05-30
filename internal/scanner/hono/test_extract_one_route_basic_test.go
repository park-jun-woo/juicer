//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_Basic 테스트
package hono

import "testing"

func TestExtractOneRoute_Basic(t *testing.T) {
	r := oneRoute(t, `app.get("/users", handler);`, map[string]bool{"app": true})
	if r == nil {
		t.Fatal("nil route")
	}
	if r.Method != "GET" || r.Path != "/users" || r.Handler != "handler" || r.OwnerVar != "app" {
		t.Fatalf("got %+v", r)
	}
}
