//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRouters_FromType 테스트
package express

import "testing"

func TestCollectParamRouters_FromType(t *testing.T) {
	src := []byte(`function setup(r: Router) { r.get('/x', h); }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if !routers["r"] {
		t.Fatalf("expected router 'r' from type annotation, got %v", routers)
	}
}
