//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_PlainCall 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_PlainCall(t *testing.T) {

	src := []byte(`get('/x', h);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for plain call, got %v", routers)
	}
}
