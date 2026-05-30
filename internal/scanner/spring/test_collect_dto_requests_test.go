//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestCollectDTORequests 테스트
package spring

import "testing"

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ci := controllerInfo{imports: map[string]string{}}
	if reqs := collectDTORequests(ep, ci, "/abs", 0); len(reqs) != 2 {
		t.Fatalf("got %d", len(reqs))
	}
}
