//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAnnotatedDepends_NoDepends 테스트
package fastapi

import "testing"

func TestClassifyAnnotatedDepends_NoDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("x", "Annotated[int, Query()]", nil, ri)
	if len(ri.middleware) != 0 || ri.bodyType != "" {
		t.Fatalf("expected no-op, got %+v", ri)
	}
}
