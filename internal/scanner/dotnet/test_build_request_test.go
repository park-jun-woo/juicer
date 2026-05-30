//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestBuildRequest 테스트
package dotnet

import "testing"

func TestBuildRequest(t *testing.T) {
	if r := buildRequest(endpointInfo{}); r != nil {
		t.Fatalf("nil: %+v", r)
	}
	r := buildRequest(endpointInfo{bodyType: "UserDto"})
	if r == nil || r.Body == nil || r.Body.Method != "FromBody" {
		t.Fatalf("body: %+v", r)
	}
}
