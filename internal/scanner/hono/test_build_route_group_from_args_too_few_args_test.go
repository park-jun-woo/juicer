//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRouteGroupFromArgs_TooFewArgs 테스트
package hono

import "testing"

func TestBuildRouteGroupFromArgs_TooFewArgs(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/users");`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("single arg should be nil, got %+v", g)
	}
}
