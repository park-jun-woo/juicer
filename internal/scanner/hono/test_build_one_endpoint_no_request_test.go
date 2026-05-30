//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildOneEndpoint_NoRequest 테스트
package hono

import "testing"

func TestBuildOneEndpoint_NoRequest(t *testing.T) {
	r := routeInfo{Method: "GET", Handler: "h"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	ep := buildOneEndpoint("GET", "/health", r, "x.ts", nil, ctx, fi)
	if ep.Request != nil {
		t.Fatalf("expected nil request, got %+v", ep.Request)
	}
}
