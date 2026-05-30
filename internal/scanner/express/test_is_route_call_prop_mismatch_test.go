//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouteCall_PropMismatch 테스트
package express

import "testing"

func TestIsRouteCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if isRouteCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}

}
