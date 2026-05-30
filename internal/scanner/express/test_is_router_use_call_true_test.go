//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsRouterUseCall_True 테스트
package express

import "testing"

func TestIsRouterUseCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`router.use('/x', r);`))
	if !isRouterUseCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected true")
	}
}
