//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_NonHttpMethod 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_NonHttpMethod(t *testing.T) {

	src := []byte(`r.foo('/x');`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for non-http method, got %v", routers)
	}
}
