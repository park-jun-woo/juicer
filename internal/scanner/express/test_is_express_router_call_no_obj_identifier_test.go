//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressRouterCall_NoObjIdentifier 테스트
package express

import "testing"

func TestIsExpressRouterCall_NoObjIdentifier(t *testing.T) {

	fi := mustParse(t, []byte(`const r = a.b.Router();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false")
	}
}
