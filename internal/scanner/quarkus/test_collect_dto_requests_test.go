//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestCollectDTORequests 테스트
package quarkus

import "testing"

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ri := resourceInfo{imports: map[string]string{}, absFile: "/abs/R.java"}
	reqs := collectDTORequests(ep, ri, "/abs", 0)
	if len(reqs) != 2 {
		t.Fatalf("expected 2 reqs, got %d", len(reqs))
	}
}
