//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectDTORequests_Empty 테스트
package nestjs

import "testing"

func TestCollectDTORequests_Empty(t *testing.T) {
	ep := endpointInfo{}
	reqs := collectDTORequests(ep, nil, "/src/controller.ts", 0)
	if len(reqs) != 0 {
		t.Fatalf("expected 0, got %d", len(reqs))
	}
}
