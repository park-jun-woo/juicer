//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildResponse 테스트
package spring

import "testing"

func TestBuildResponse(t *testing.T) {
	if r := buildResponse(endpointInfo{}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
	r := buildResponse(endpointInfo{method: "POST", returnType: "UserDto"})
	if r == nil || r.Status != "201" || r.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}
