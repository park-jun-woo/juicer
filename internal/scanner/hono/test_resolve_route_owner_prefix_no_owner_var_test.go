//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveRouteOwnerPrefix_NoOwnerVar 테스트
package hono

import "testing"

func TestResolveRouteOwnerPrefix_NoOwnerVar(t *testing.T) {
	ctx := &scanContext{prefixMap: map[string]string{
		prefixKey("f.ts", ""): "/x",
	}}
	r := routeInfo{OwnerVar: ""}
	if got := resolveRouteOwnerPrefix(r, ctx, "f.ts"); got != "" {
		t.Fatalf("expected empty for no owner var, got %q", got)
	}
}
