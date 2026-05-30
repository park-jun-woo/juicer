//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAnnotatedDepends_Middleware 테스트
package fastapi

import "testing"

func TestClassifyAnnotatedDepends_Middleware(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("dep", "Annotated[User, Depends(get_current_user)]", nil, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_current_user" {
		t.Fatalf("got %+v", ri.middleware)
	}
}
