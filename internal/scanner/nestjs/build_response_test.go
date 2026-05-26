//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildResponse_CustomCode 테스트
package nestjs

import "testing"

func TestBuildResponse_CustomCode(t *testing.T) {
	ep := endpointInfo{method: "POST", statusCode: 201, returnType: "UserDto"}
	resp := buildResponse(ep)
	if resp.Status != "201" {
		t.Fatalf("expected 201, got %s", resp.Status)
	}
}
