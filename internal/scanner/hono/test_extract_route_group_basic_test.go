//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_Basic 테스트
package hono

import "testing"

func TestExtractRouteGroup_Basic(t *testing.T) {
	g := oneRouteGroup(t, `app.route("/api", subApp);`, map[string]bool{"app": true, "subApp": true})
	if g == nil || g.Prefix != "/api" || g.ParentVar != "app" || g.SubAppName != "subApp" {
		t.Fatalf("got %+v", g)
	}
}
