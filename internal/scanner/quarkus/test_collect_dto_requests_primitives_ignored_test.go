//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestCollectDTORequests_PrimitivesIgnored 테스트
package quarkus

import "testing"

func TestCollectDTORequests_PrimitivesIgnored(t *testing.T) {
	ep := endpointInfo{bodyType: "String", returnType: "int"}
	ri := resourceInfo{imports: map[string]string{}}
	if reqs := collectDTORequests(ep, ri, "/abs", 0); len(reqs) != 0 {
		t.Fatalf("expected 0, got %d", len(reqs))
	}
}
