//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestBuildRequest_Body 테스트
package spring

import "testing"

func TestBuildRequest_Body(t *testing.T) {
	r := buildRequest(endpointInfo{bodyType: "UserDto"})
	if r == nil || r.Body == nil || r.Body.Method != "RequestBody" {
		t.Fatalf("got %+v", r)
	}
}
