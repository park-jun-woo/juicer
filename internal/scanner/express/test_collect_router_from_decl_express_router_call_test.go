//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterFromDecl_ExpressRouterCall 테스트
package express

import "testing"

func TestCollectRouterFromDecl_ExpressRouterCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if !routers["r"] {
		t.Fatalf("expected r, got %v", routers)
	}
}
