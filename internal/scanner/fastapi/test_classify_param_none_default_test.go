//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_NoneDefault 테스트
package fastapi

import "testing"

func TestClassifyParam_NoneDefault(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(q: str = None): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || !ri.query[0].DefaultIsNull {
		t.Fatalf("expected nullable query, got %+v", ri.query)
	}
}
