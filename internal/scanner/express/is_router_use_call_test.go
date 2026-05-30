//ff:func feature=scan type=test control=sequence topic=express
//ff:what isRouterUseCall: router.use true / 비멤버 / 미등록 / prop불일치
package express

import "testing"

func TestIsRouterUseCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`router.use('/x', r);`))
	if !isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected true")
	}
}

func TestIsRouterUseCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`use('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouterUseCall_Unregistered(t *testing.T) {
	fi := mustParse(t, []byte(`other.use('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouterUseCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}
