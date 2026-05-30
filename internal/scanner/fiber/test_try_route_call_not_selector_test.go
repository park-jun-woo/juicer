//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NotSelector 테스트
package fiber

import "testing"

func TestTryRouteCall_NotSelector(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { plainCall() }\n")
	if call == nil {

		return
	}
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{}, "m.go", fset)
	if ok {
		t.Fatal("plain call should be false")
	}
}
