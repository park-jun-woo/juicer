//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestRouterVarOfCall_NoMember 테스트
package express

import "testing"

func TestRouterVarOfCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
