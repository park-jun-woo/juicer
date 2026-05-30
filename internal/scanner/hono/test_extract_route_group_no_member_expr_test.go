//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_NoMemberExpr 테스트
package hono

import "testing"

func TestExtractRouteGroup_NoMemberExpr(t *testing.T) {
	if g := oneRouteGroup(t, `foo();`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}
