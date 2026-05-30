//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what classifyAsMiddleware: alias / inline Annotated / Depends없음 분기
package fastapi

import "testing"

func TestClassifyAsMiddleware_Alias(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("AuthDep", map[string]string{"AuthDep": "verify"}, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "verify" {
		t.Fatalf("got %+v", ri.middleware)
	}
}

func TestClassifyAsMiddleware_Inline(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("Annotated[User, Depends(get_user)]", nil, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_user" {
		t.Fatalf("got %+v", ri.middleware)
	}
}

func TestClassifyAsMiddleware_NoDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("int", nil, ri)
	if len(ri.middleware) != 0 {
		t.Fatalf("expected no-op, got %+v", ri.middleware)
	}
}
