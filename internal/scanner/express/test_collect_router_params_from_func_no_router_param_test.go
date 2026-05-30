//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterParamsFromFunc_NoRouterParam 테스트
package express

import "testing"

func TestCollectRouterParamsFromFunc_NoRouterParam(t *testing.T) {
	fi := mustParse(t, []byte(`function setup(n: number) {}`))
	routers := map[string]bool{}
	collectRouterParamsFromFunc(firstFuncDecl(t, fi), fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
