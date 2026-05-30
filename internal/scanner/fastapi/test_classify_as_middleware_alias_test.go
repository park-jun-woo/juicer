//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAsMiddleware_Alias 테스트
package fastapi

import "testing"

func TestClassifyAsMiddleware_Alias(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("AuthDep", map[string]string{"AuthDep": "verify"}, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "verify" {
		t.Fatalf("got %+v", ri.middleware)
	}
}
