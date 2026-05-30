//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterUseCall_PropMismatch 테스트
package express

import "testing"

func TestIsRouterUseCall_PropMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`router.get('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}
