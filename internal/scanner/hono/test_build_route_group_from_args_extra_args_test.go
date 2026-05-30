//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRouteGroupFromArgs_ExtraArgs 테스트
package hono

import "testing"

func TestBuildRouteGroupFromArgs_ExtraArgs(t *testing.T) {

	call, src := firstCallExpr(t, `app.route("/x", sub, extra);`+"\n")
	g := buildRouteGroupFromArgs(call, src, "app")
	if g == nil || g.Prefix != "/x" || g.SubAppName != "sub" {
		t.Fatalf("group = %+v", g)
	}
}
