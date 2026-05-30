//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestBuildEndpoint_NoRequest 테스트
package flask

import "testing"

func TestBuildEndpoint_NoRequest(t *testing.T) {
	ri := routeInfo{method: "GET", path: "/health", handler: "health", file: "app.py", line: 1}
	ep := buildEndpoint(ri)
	if ep.Request != nil {
		t.Fatalf("expected no request, got %+v", ep.Request)
	}
}
