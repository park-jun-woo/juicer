//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAsMiddleware_NoDepends 테스트
package fastapi

import "testing"

func TestClassifyAsMiddleware_NoDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("int", nil, ri)
	if len(ri.middleware) != 0 {
		t.Fatalf("expected no-op, got %+v", ri.middleware)
	}
}
