//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBodyContainsRouterUse_Round5 테스트
package express

import "testing"

func TestBodyContainsRouterUse_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const f = () => { router.use('/x', sub); };`))
	arrow := exFirst(t, fi, "arrow_function")
	body := extractFunctionBody(arrow)
	if !bodyContainsRouterUse(body, fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected router.use detected")
	}
	fi2 := mustParse(t, []byte(`const g = () => { return 1; };`))
	arrow2 := exFirst(t, fi2, "arrow_function")
	body2 := extractFunctionBody(arrow2)
	if bodyContainsRouterUse(body2, fi2.Src, map[string]bool{"router": true}) {
		t.Fatal("expected no router.use")
	}
}
