//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterParamsFromFunc_WithRouterParam 테스트
package express

import "testing"

func TestCollectRouterParamsFromFunc_WithRouterParam(t *testing.T) {
	fi := mustParse(t, []byte(`function setup(r: Router) {}`))
	routers := map[string]bool{}
	collectRouterParamsFromFunc(firstFuncDecl(t, fi), fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r, got %v", routers)
	}
}
