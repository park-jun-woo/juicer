//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRouteGroupFromArgs_PrefixNotString 테스트
package hono

import "testing"

func TestBuildRouteGroupFromArgs_PrefixNotString(t *testing.T) {
	call, src := firstCallExpr(t, `app.route(prefixVar, sub);`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("non-string prefix should be nil, got %+v", g)
	}
}
