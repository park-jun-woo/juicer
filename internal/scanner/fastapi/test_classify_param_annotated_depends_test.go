//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_AnnotatedDepends 테스트
package fastapi

import "testing"

func TestClassifyParam_AnnotatedDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(user: Annotated[User, Depends(get_user)]): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_user" {
		t.Fatalf("expected middleware, got %+v", ri.middleware)
	}
}
