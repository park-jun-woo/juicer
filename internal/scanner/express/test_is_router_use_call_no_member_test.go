//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterUseCall_NoMember 테스트
package express

import "testing"

func TestIsRouterUseCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`use('/x');`))
	if isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected false")
	}
}
