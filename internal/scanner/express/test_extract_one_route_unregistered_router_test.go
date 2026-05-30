//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractOneRoute_UnregisteredRouter 테스트
package express

import "testing"

func TestExtractOneRoute_UnregisteredRouter(t *testing.T) {
	fi := mustParse(t, []byte(`other.get('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for unregistered router, got %+v", ri)
	}
}
