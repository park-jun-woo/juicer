//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestBuildResponse_HTML 테스트
package fastapi

import "testing"

func TestBuildResponse_HTML(t *testing.T) {
	resp := buildResponse(routeInfo{method: "GET", responseClass: "HTMLResponse"}, "")
	if resp.Kind != "html" {
		t.Fatalf("expected html kind, got %+v", resp)
	}
}
