//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressRouterCall_NoMember 테스트
package express

import "testing"

func TestIsExpressRouterCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express();`))
	if isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected false for plain call")
	}
}
