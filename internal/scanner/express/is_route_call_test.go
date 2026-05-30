//ff:func feature=scan type=test control=sequence topic=express
//ff:what isRouteCall: router.route() true / 비call / 비멤버 / 비obj / 미등록 / prop불일치
package express

import "testing"

func TestIsRouteCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/x');`))
	if !isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected true")
	}
}

func TestIsRouteCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`router;`))
	ids := findAllByType(fi.Root, "identifier")
	if isRouteCall(ids[0], fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouteCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`route('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouteCall_NoObjIdent(t *testing.T) {
	fi := mustParse(t, []byte(`a.b.route('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"a": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouteCall_Unregistered(t *testing.T) {
	fi := mustParse(t, []byte(`other.route('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}

func TestIsRouteCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
	// remaining uncovered branch (prop==nil) is unreachable for valid member_expression
}
