//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_NotHonoVar 테스트
package hono

import "testing"

func TestExtractRouteGroup_NotHonoVar(t *testing.T) {
	if g := oneRouteGroup(t, `app.route("/x", s);`, map[string]bool{"other": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}
