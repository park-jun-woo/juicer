//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildResponse 테스트
package dotnet

import "testing"

func TestBuildResponse(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("nil: %+v", r)
	}
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}
