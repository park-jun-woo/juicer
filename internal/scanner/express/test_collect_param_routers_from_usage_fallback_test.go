//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRouters_FromUsageFallback 테스트
package express

import "testing"

func TestCollectParamRouters_FromUsageFallback(t *testing.T) {

	src := []byte(`function setup(r) { r.get('/x', h); }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if !routers["r"] {
		t.Fatalf("expected router 'r' from usage fallback, got %v", routers)
	}
}
