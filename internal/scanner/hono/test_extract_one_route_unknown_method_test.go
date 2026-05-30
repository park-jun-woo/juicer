//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_UnknownMethod 테스트
package hono

import "testing"

func TestExtractOneRoute_UnknownMethod(t *testing.T) {
	if r := oneRoute(t, `app.use("/x", h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
