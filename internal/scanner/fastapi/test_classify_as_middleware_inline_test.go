//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAsMiddleware_Inline 테스트
package fastapi

import "testing"

func TestClassifyAsMiddleware_Inline(t *testing.T) {
	ri := &routeInfo{}
	classifyAsMiddleware("Annotated[User, Depends(get_user)]", nil, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_user" {
		t.Fatalf("got %+v", ri.middleware)
	}
}
