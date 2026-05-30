//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRequest_NoContent 테스트
package hono

import "testing"

func TestBuildRequest_NoContent(t *testing.T) {
	r := routeInfo{Method: "GET"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	if req := buildRequest(r, nil, ctx, fi); req != nil {
		t.Fatalf("expected nil for no content, got %+v", req)
	}
}
