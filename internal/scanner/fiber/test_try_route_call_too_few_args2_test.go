//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_TooFewArgs2 테스트
package fiber

import "testing"

func TestTryRouteCall_TooFewArgs2(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { app.Get(\"/x\") }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("single arg should be false")
	}
}
