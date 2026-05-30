//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_ExpressIgnored 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_ExpressIgnored(t *testing.T) {
	src := []byte(`express.get('/x', h); module.get('/y'); exports.get('/z');`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected express/module/exports ignored, got %v", routers)
	}
}
