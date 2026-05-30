//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyParam_PathParam 테스트
package fastapi

import "testing"

func TestClassifyParam_PathParam(t *testing.T) {
	ri := &routeInfo{}
	classifyAllParams([]byte("def f(item_id: int): pass\n"), ri, map[string]bool{"item_id": true}, nil)
	if len(ri.params) != 1 || ri.params[0].Name != "item_id" {
		t.Fatalf("expected path param, got %+v", ri.params)
	}
}
