//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestCollectDTORequests_Primitives 테스트
package dotnet

import "testing"

func TestCollectDTORequests_Primitives(t *testing.T) {
	ep := endpointInfo{bodyType: "string", returnType: "int"}
	if reqs := collectDTORequests(ep, controllerInfo{}, "/abs", 0); len(reqs) != 0 {
		t.Fatalf("got %d", len(reqs))
	}
}
