//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractRouteGroup_SingleQuotedPrefix 테스트
package hono

import "testing"

func TestExtractRouteGroup_SingleQuotedPrefix(t *testing.T) {
	g := oneRouteGroup(t, `app.route('/v2', api);`, map[string]bool{"app": true})
	if g == nil || g.Prefix != "/v2" || g.SubAppName != "api" {
		t.Fatalf("got %+v", g)
	}
}
