//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterStandaloneCall_MemberCallee 테스트
package express

import "testing"

func TestIsRouterStandaloneCall_MemberCallee(t *testing.T) {

	fi := mustParse(t, []byte(`a.b();`))
	if isRouterStandaloneCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for member callee")
	}
}
