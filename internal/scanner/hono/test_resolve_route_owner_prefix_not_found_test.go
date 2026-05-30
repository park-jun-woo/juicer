//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRouteOwnerPrefix_NotFound 테스트
package hono

import "testing"

func TestResolveRouteOwnerPrefix_NotFound(t *testing.T) {
	ctx := &scanContext{prefixMap: map[string]string{}}
	r := routeInfo{OwnerVar: "app"}
	if got := resolveRouteOwnerPrefix(r, ctx, "f.ts"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
