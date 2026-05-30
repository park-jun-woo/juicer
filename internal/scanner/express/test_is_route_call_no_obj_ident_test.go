//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouteCall_NoObjIdent 테스트
package express

import "testing"

func TestIsRouteCall_NoObjIdent(t *testing.T) {
	fi := mustParse(t, []byte(`a.b.route('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"a": true}) {
		t.Fatal("expected false")
	}
}
