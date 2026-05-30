//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindHandlerBody_FromHandlerNode 테스트
package express

import "testing"

func TestFindHandlerBody_FromHandlerNode(t *testing.T) {
	fi := mustParse(t, []byte(`const h = (req, res) => { res.json({}); };`))
	ri := routeInfo{HandlerNode: firstArrow(t, fi)}
	if body := findHandlerBody(fi, ri); body == nil {
		t.Fatal("expected body from handler node")
	}
}
