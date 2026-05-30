//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildControllerEndpoints_Round5 테스트
package spring

import "testing"

func TestBuildControllerEndpoints_Round5(t *testing.T) {
	ci := controllerInfo{
		prefix:    "/users",
		className: "UserController",
		file:      "C.java",
		endpoints: []endpointInfo{
			{method: "GET", path: "/{id}", handler: "get", returnType: "UserDto"},
		},
		imports: map[string]string{},
	}
	eps, _ := buildControllerEndpoints(ci, "/abs", 0)
	if len(eps) != 1 || eps[0].Path != "/users/{id}" {
		t.Fatalf("endpoints: %+v", eps)
	}
}
