//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRequest_PathParams 테스트
package hono

import "testing"

func TestBuildRequest_PathParams(t *testing.T) {
	r := routeInfo{Method: "GET"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	req := buildRequest(r, []string{"id", "slug"}, ctx, fi)
	if req == nil || len(req.PathParams) != 2 {
		t.Fatalf("expected 2 path params, got %+v", req)
	}
}
