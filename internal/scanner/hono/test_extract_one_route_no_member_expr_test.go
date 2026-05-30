//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_NoMemberExpr 테스트
package hono

import "testing"

func TestExtractOneRoute_NoMemberExpr(t *testing.T) {
	if r := oneRoute(t, `foo();`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
