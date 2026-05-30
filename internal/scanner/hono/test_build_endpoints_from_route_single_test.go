//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildEndpointsFromRoute_Single 테스트
package hono

import "testing"

func TestBuildEndpointsFromRoute_Single(t *testing.T) {
	r := routeInfo{Method: "GET", Path: "/users/:id", Handler: "h", OwnerVar: "app"}
	ctx := minimalCtx()
	ctx.prefixMap[prefixKey("/abs/app.ts", "app")] = "/api"
	fi := &fileInfo{Path: "/abs/app.ts"}

	eps := buildEndpointsFromRoute(r, map[string]bool{}, ctx, fi, "app.ts", "/abs/app.ts")
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Method != "GET" || eps[0].Path != "/api/users/{id}" {
		t.Fatalf("ep = %+v", eps[0])
	}
}
