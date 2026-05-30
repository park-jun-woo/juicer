//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRouteGroupFromArgs_SubAppNotIdent 테스트
package hono

import "testing"

func TestBuildRouteGroupFromArgs_SubAppNotIdent(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/x", "notident");`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("non-ident subapp should be nil, got %+v", g)
	}
}
