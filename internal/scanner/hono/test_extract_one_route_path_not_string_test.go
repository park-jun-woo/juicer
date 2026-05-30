//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneRoute_PathNotString 테스트
package hono

import "testing"

func TestExtractOneRoute_PathNotString(t *testing.T) {
	if r := oneRoute(t, `app.get(pathVar, h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
