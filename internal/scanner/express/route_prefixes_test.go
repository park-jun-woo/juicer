//ff:func feature=scan type=test control=sequence topic=express
//ff:what routePrefixes: 등록된 prefix 반환 / 없으면 빈 prefix 하나
package express

import "testing"

func TestRoutePrefixes_Found(t *testing.T) {
	ctx := &scanContext{routerPrefixes: map[routerKey][]string{
		{file: "a.ts", varName: "r"}: {"/api", "/v2"},
	}}
	got := routePrefixes(ctx, "a.ts", "r")
	if len(got) != 2 || got[0] != "/api" {
		t.Fatalf("got %v", got)
	}
}

func TestRoutePrefixes_Default(t *testing.T) {
	ctx := &scanContext{routerPrefixes: map[routerKey][]string{}}
	got := routePrefixes(ctx, "a.ts", "r")
	if len(got) != 1 || got[0] != "" {
		t.Fatalf("got %v", got)
	}
}
