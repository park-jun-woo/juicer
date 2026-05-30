//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressCall_True 테스트
package express

import "testing"

func TestIsExpressCall_True(t *testing.T) {
	fi := mustParse(t, []byte(`const a = express();`))
	if !isExpressCall(firstCallExpr(t, fi), fi.Src) {
		t.Fatal("expected true")
	}
}
