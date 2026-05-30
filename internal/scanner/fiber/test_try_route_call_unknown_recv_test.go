//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_UnknownRecv 테스트
package fiber

import "testing"

func TestTryRouteCall_UnknownRecv(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { other.Get(\"/x\", h) }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("unknown receiver should be false")
	}
}
