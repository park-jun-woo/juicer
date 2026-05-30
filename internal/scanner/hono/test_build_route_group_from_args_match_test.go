//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRouteGroupFromArgs_Match 테스트
package hono

import "testing"

func TestBuildRouteGroupFromArgs_Match(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/users", usersApp);`+"\n")
	g := buildRouteGroupFromArgs(call, src, "app")
	if g == nil || g.Prefix != "/users" || g.SubAppName != "usersApp" || g.ParentVar != "app" {
		t.Fatalf("group = %+v", g)
	}
}
