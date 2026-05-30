//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromUsage_NoArgs 테스트
package express

import "testing"

func TestCollectParamRoutersFromUsage_NoArgs(t *testing.T) {
	src := []byte(`r.get();`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromUsage(fi, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none for empty args, got %v", routers)
	}
}
