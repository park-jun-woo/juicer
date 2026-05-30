//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_NoIdentifierObject 테스트
package hono

import "testing"

func TestExtractRouteGroup_NoIdentifierObject(t *testing.T) {
	if g := oneRouteGroup(t, `this.route("/x", s);`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}
