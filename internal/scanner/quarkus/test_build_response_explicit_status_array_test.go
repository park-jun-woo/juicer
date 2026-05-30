//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildResponse_ExplicitStatusArray 테스트
package quarkus

import "testing"

func TestBuildResponse_ExplicitStatusArray(t *testing.T) {
	r := buildResponse(endpointInfo{method: "GET", statusCode: "200", returnType: "UserDto", returnIsArray: true})
	if r == nil || r.Status != "200" || r.Body != "array" {
		t.Fatalf("got %+v", r)
	}
}
