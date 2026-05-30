//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_DefaultValueQuery 테스트
package fastapi

import "testing"

func TestClassifyParam_DefaultValueQuery(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(limit: int = 10): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 1 || ri.query[0].Default != "10" {
		t.Fatalf("expected query with default, got %+v", ri.query)
	}
}
