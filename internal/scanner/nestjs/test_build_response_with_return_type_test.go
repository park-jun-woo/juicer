//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildResponse_WithReturnType 테스트
package nestjs

import "testing"

func TestBuildResponse_WithReturnType(t *testing.T) {
	ep := endpointInfo{method: "GET", returnType: "UserDto"}
	resp := buildResponse(ep)
	if resp.TypeName != "UserDto" {
		t.Fatalf("expected UserDto, got %q", resp.TypeName)
	}
}
