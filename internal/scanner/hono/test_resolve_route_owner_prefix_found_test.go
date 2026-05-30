//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRouteOwnerPrefix_Found 테스트
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
