//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_TypedQuery 테스트
package fastapi

import "testing"

func TestClassifyParam_TypedQuery(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(tag: str): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || ri.query[0].Name != "tag" {
		t.Fatalf("expected typed query, got %+v", ri.query)
	}
}
