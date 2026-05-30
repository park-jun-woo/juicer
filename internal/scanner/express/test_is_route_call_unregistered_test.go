//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouteCall_Unregistered 테스트
package express

import "testing"

func TestIsRouteCall_Unregistered(t *testing.T) {
	fi := mustParse(t, []byte(`other.route('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}
