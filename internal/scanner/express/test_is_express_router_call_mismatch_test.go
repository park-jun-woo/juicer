//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressRouterCall_Mismatch 테스트
package express

import "testing"

func TestIsExpressRouterCall_Mismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = foo.Bar();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}
