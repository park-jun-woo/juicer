//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_NonStringFirstArg 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_NonStringFirstArg(t *testing.T) {

	src := []byte(`r.use(mw);`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for non-string first arg, got %v", routers)
	}
}
