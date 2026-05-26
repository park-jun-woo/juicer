//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestBuildResponse_DefaultCode 테스트
package nestjs

import "testing"

func TestBuildResponse_DefaultCode(t *testing.T) {
	ep := endpointInfo{method: "GET"}
	resp := buildResponse(ep)
	if resp.Status != "200" {
		t.Fatalf("expected 200, got %s", resp.Status)
	}
}
