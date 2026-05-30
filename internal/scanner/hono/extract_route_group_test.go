//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractRouteGroup 테스트
package hono

import "testing"

func oneRouteGroup(t *testing.T, src string, vars map[string]bool) *routeGroup {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	return extractRouteGroup(call, fi.Src, vars)
}

func TestExtractRouteGroup_Basic(t *testing.T) {
	g := oneRouteGroup(t, `app.route("/api", subApp);`, map[string]bool{"app": true, "subApp": true})
	if g == nil || g.Prefix != "/api" || g.ParentVar != "app" || g.SubAppName != "subApp" {
		t.Fatalf("got %+v", g)
	}
}

func TestExtractRouteGroup_NoMemberExpr(t *testing.T) {
	if g := oneRouteGroup(t, `foo();`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}

func TestExtractRouteGroup_NoIdentifierObject(t *testing.T) {
	if g := oneRouteGroup(t, `this.route("/x", s);`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}

func TestExtractRouteGroup_NotHonoVar(t *testing.T) {
	if g := oneRouteGroup(t, `app.route("/x", s);`, map[string]bool{"other": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}

func TestExtractRouteGroup_NotRouteMethod(t *testing.T) {
	if g := oneRouteGroup(t, `app.get("/x", h);`, map[string]bool{"app": true}); g != nil {
		t.Fatalf("expected nil, got %+v", g)
	}
}

func TestExtractRouteGroup_SingleQuotedPrefix(t *testing.T) {
	g := oneRouteGroup(t, `app.route('/v2', api);`, map[string]bool{"app": true})
	if g == nil || g.Prefix != "/v2" || g.SubAppName != "api" {
		t.Fatalf("got %+v", g)
	}
}
