//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildResponse_DefaultStatus 테스트
package quarkus

import "testing"

func TestBuildResponse_DefaultStatus(t *testing.T) {
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}
