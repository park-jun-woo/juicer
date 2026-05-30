//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildOneEndpoint_WithPathParams 테스트
package hono

import "testing"

func TestBuildOneEndpoint_WithPathParams(t *testing.T) {
	r := routeInfo{Method: "GET", Handler: "h", Line: 3, Middleware: []string{"auth"}}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	ep := buildOneEndpoint("GET", "/users/{id}", r, "x.ts", []string{"id"}, ctx, fi)
	if ep.Method != "GET" || ep.Path != "/users/{id}" || ep.Line != 3 {
		t.Fatalf("ep = %+v", ep)
	}
	if len(ep.Middleware) != 1 {
		t.Errorf("middleware not propagated: %v", ep.Middleware)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected path params, got %+v", ep.Request)
	}
}
