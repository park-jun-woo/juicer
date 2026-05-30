//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_AllMethod 테스트
package hono

import "testing"

func TestExtractOneRoute_AllMethod(t *testing.T) {
	r := oneRoute(t, `app.all("/x", handler);`, map[string]bool{"app": true})
	if r == nil || r.Method != "all" {
		t.Fatalf("got %+v", r)
	}
}
