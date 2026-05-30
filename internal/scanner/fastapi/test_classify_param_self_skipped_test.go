//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_SelfSkipped 테스트
package fastapi

import "testing"

func TestClassifyParam_SelfSkipped(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(self, cls): pass\n"), ri, map[string]bool{}, nil)
	if len(ri.query) != 0 || len(ri.params) != 0 {
		t.Fatalf("self/cls should be skipped, got %+v", ri)
	}
}
