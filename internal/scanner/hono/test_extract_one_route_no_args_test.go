//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_NoArgs 테스트
package hono

import "testing"

func TestExtractOneRoute_NoArgs(t *testing.T) {
	if r := oneRoute(t, `app.get();`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
