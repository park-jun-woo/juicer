//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectRouterFromDecl_NoInitValue 테스트
package express

import "testing"

func TestCollectRouterFromDecl_NoInitValue(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 5;`))
	routers := map[string]bool{}
	collectRouterFromDecl(firstLexDecl(t, fi), fi, routers, map[string]bool{})
	if len(routers) != 0 {
		t.Fatalf("expected none, got %v", routers)
	}
}
