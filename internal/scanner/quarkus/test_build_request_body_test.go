//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestBuildRequest_Body 테스트
package quarkus

import "testing"

func TestBuildRequest_Body(t *testing.T) {
	r := buildRequest(endpointInfo{bodyType: "UserDto", bodyVarName: "dto"})
	if r == nil || r.Body == nil || r.Body.TypeName != "UserDto" {
		t.Fatalf("got %+v", r)
	}
}
