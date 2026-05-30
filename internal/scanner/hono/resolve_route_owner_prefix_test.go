//ff:func feature=scan type=test control=sequence topic=hono
//ff:what resolveRouteOwnerPrefix 테스트
package hono

import "testing"

func TestResolveRouteOwnerPrefix_Found(t *testing.T) {
	ctx := &scanContext{prefixMap: map[string]string{
		prefixKey("f.ts", "app"): "/api",
	}}
	r := routeInfo{OwnerVar: "app"}
	if got := resolveRouteOwnerPrefix(r, ctx, "f.ts"); got != "/api" {
		t.Fatalf("got %q", got)
	}
}

func TestResolveRouteOwnerPrefix_NotFound(t *testing.T) {
	ctx := &scanContext{prefixMap: map[string]string{}}
	r := routeInfo{OwnerVar: "app"}
	if got := resolveRouteOwnerPrefix(r, ctx, "f.ts"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveRouteOwnerPrefix_NoOwnerVar(t *testing.T) {
	ctx := &scanContext{prefixMap: map[string]string{
		prefixKey("f.ts", ""): "/x",
	}}
	r := routeInfo{OwnerVar: ""}
	if got := resolveRouteOwnerPrefix(r, ctx, "f.ts"); got != "" {
		t.Fatalf("expected empty for no owner var, got %q", got)
	}
}
