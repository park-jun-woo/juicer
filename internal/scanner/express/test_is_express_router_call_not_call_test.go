//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestIsExpressRouterCall_NotCall 테스트
package express

import "testing"

func TestIsExpressRouterCall_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express;`))
	ids := findAllByType(fi.Root, "identifier")
	if isExpressRouterCall(ids[0], fi.Src) {
		t.Fatal("expected false for non-call")
	}
}
