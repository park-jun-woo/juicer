//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectParamRoutersFromType_None 테스트
package express

import "testing"

func TestCollectParamRoutersFromType_None(t *testing.T) {
	src := []byte(`function fd(a: number) {}`)
	fi := mustParse(t, src)
	routers := map[string]bool{}
	collectParamRoutersFromType(fi, routers)
	if len(routers) != 0 {
		t.Errorf("expected none, got %v", routers)
	}
}
