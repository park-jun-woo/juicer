//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildRequest_WithBody 테스트
package nestjs

import "testing"

func TestBuildRequest_WithBody(t *testing.T) {
	ep := endpointInfo{bodyType: "CreateUserDto"}
	req := buildRequest(ep)
	if req == nil || req.Body == nil {
		t.Fatal("expected body")
	}
	if req.Body.TypeName != "CreateUserDto" {
		t.Fatalf("expected CreateUserDto, got %q", req.Body.TypeName)
	}
}
