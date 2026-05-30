//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestRouterVarOfCall_NoIdentifier 테스트
package express

import "testing"

func TestRouterVarOfCall_NoIdentifier(t *testing.T) {
	fi := mustParse(t, []byte(`a.b.get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
