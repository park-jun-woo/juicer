//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_NoIdentifierObject 테스트
package hono

import "testing"

func TestExtractOneRoute_NoIdentifierObject(t *testing.T) {
	if r := oneRoute(t, `this.get("/x", h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
