//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildResourceEndpoints_Round5 테스트
package quarkus

import "testing"

func TestBuildResourceEndpoints_Round5(t *testing.T) {
	ri := resourceInfo{
		prefix: "/users",
		endpoints: []endpointInfo{
			{method: "GET", path: "/{id}", handler: "get", returnType: "UserDto"},
		},
	}
	eps, _ := buildResourceEndpoints(ri, "/abs", 0)
	if len(eps) != 1 || eps[0].Path != "/users/{id}" {
		t.Fatalf("endpoints: %+v", eps)
	}
}
