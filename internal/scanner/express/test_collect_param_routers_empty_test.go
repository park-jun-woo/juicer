//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRouters_Empty 테스트
package express

import "testing"

func TestCollectParamRouters_Empty(t *testing.T) {
	src := []byte(`function setup() { return 1; }`)
	fi := mustParse(t, src)
	routers := collectParamRouters(fi)
	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}
