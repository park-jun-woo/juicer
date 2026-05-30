//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_NotRouteMethod 테스트
package hono

import "testing"

func TestExtractRouteGroup_NotRouteMethod(t *testing.T) {
	if g := oneRouteGroup(t, `app.get("/x", h);`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}
