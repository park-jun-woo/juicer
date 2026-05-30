//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterUseCall_Unregistered 테스트
package express

import "testing"

func TestIsRouterUseCall_Unregistered(t *testing.T) {
	fi := mustParse(t, []byte(`other.use('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}
