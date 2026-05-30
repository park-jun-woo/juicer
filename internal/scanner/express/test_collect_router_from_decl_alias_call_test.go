//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterFromDecl_AliasCall 테스트
package express

import "testing"

func TestCollectRouterFromDecl_AliasCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{"Router": true})
	if !routers["r"] {
		t.Fatalf("expected r via alias, got %v", routers)
	}
}
