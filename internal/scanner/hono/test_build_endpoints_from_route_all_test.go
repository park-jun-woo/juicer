//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildEndpointsFromRoute_All 테스트
package hono

import "testing"

func TestBuildEndpointsFromRoute_All(t *testing.T) {
	r := routeInfo{Method: "all", Path: "/", Handler: "h"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	eps := buildEndpointsFromRoute(r, map[string]bool{}, ctx, fi, "x.ts", "/abs/x.ts")
	if len(eps) != 5 {
		t.Fatalf("expected 5 endpoints for 'all', got %d", len(eps))
	}
}
