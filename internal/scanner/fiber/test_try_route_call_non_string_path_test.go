//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NonStringPath 테스트
package fiber

import "testing"

func TestTryRouteCall_NonStringPath(t *testing.T) {
	call, fset := routeCallFrom(t, "package m\nfunc f() { app.Get(pathVar, h) }\n")
	_, _, ok := tryRouteCall(call, map[string]*routerInfo{"app": {}}, "m.go", fset)
	if ok {
		t.Fatal("non-string path should be false")
	}
}
