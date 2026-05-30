//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestRoutePrefixes_Found 테스트
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
