//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterFromDecl_NotRouterCall 테스트
package express

import "testing"

func TestCollectRouterFromDecl_NotRouterCall(t *testing.T) {
	fi := mustParse(t, []byte(`const x = foo();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
