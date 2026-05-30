//ff:func feature=scan type=test control=sequence topic=express
//ff:what routerVarOfCall: 변수명 반환 / 비멤버 / 비식별자 분기
package express

import "testing"

func TestRouterVarOfCall_Found(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "router" {
		t.Fatalf("got %q", got)
	}
}

func TestRouterVarOfCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestRouterVarOfCall_NoIdentifier(t *testing.T) {
	fi := mustParse(t, []byte(`a.b.get('/x');`))
	if got := routerVarOfCall(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
