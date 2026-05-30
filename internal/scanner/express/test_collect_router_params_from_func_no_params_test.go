//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterParamsFromFunc_NoParams 테스트
package express

import "testing"

func TestCollectRouterParamsFromFunc_NoParams(t *testing.T) {

	fi := mustParse(t, []byte(`const f = x => x;`))
	arrows := findAllByType(fi.Root, "arrow_function")
	if len(arrows) == 0 {
		t.Fatal("no arrow_function")
	}
	routers := map[string]bool{}
	collectRouterParamsFromFunc(arrows[0], fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
