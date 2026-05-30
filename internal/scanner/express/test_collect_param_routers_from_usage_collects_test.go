//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_Collects 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_Collects(t *testing.T) {
	src := []byte(`r.get('/x', h);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if !routers["r"] {
		t.Fatalf("expected 'r', got %v", routers)
	}
}
