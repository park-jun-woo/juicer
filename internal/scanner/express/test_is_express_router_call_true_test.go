//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressRouterCall_True 테스트
package express

import "testing"

func TestIsExpressRouterCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	if !isExpressRouterCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}
