//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_NotHonoVar 테스트
package hono

import "testing"

func TestExtractOneRoute_NotHonoVar(t *testing.T) {
	if r := oneRoute(t, `app.get("/x", h);`, map[string]bool{"other": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
