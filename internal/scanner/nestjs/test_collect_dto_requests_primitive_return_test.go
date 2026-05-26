//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectDTORequests_PrimitiveReturn 테스트
package nestjs

import "testing"

func TestCollectDTORequests_PrimitiveReturn(t *testing.T) {
	ep := endpointInfo{returnType: "string"}
	reqs := collectDTORequests(ep, nil, "/src/controller.ts", "", 0)
	if len(reqs) != 0 {
		t.Fatalf("expected 0 for primitive return, got %d", len(reqs))
	}
}
