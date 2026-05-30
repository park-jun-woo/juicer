//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestRouterVarOfCall_Found 테스트
package express

import "testing"

func TestRouterVarOfCall_Found(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "router" {
		t.Fatalf("got %q", got)
	}
}
