//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestRoutePrefixes_Default 테스트
package express

import "testing"

func TestRoutePrefixes_Default(t *testing.T) {
	ctx := &scanContext{routerPrefixes: map[routerKey][]string{}}
	got := routePrefixes(ctx, "a.ts", "r")
	if len(got) != 1 || got[0] != "" {
		t.Fatalf("got %v", got)
	}
}
