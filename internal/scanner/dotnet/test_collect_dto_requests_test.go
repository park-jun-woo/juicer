//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestCollectDTORequests 테스트
package dotnet

import "testing"

func TestCollectDTORequests(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto", returnType: "UserDto"}
	ci := controllerInfo{usings: []string{}}
	reqs := collectDTORequests(ep, ci, "/abs", 0)
	if len(reqs) != 2 {
		t.Fatalf("got %d", len(reqs))
	}
}
