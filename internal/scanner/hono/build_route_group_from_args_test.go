//ff:func feature=scan type=test control=sequence topic=hono
//ff:what buildRouteGroupFromArgs 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstCallExpr(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call_expression")
	}
	return calls[0], fi.Src
}

func TestBuildRouteGroupFromArgs_Match(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/users", usersApp);`+"\n")
	g := buildRouteGroupFromArgs(call, src, "app")
	if g == nil || g.Prefix != "/users" || g.SubAppName != "usersApp" || g.ParentVar != "app" {
		t.Fatalf("group = %+v", g)
	}
}

func TestBuildRouteGroupFromArgs_TooFewArgs(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/users");`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("single arg should be nil, got %+v", g)
	}
}

func TestBuildRouteGroupFromArgs_PrefixNotString(t *testing.T) {
	call, src := firstCallExpr(t, `app.route(prefixVar, sub);`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("non-string prefix should be nil, got %+v", g)
	}
}

func TestBuildRouteGroupFromArgs_ExtraArgs(t *testing.T) {
	// more than 2 args -> still uses first two
	call, src := firstCallExpr(t, `app.route("/x", sub, extra);`+"\n")
	g := buildRouteGroupFromArgs(call, src, "app")
	if g == nil || g.Prefix != "/x" || g.SubAppName != "sub" {
		t.Fatalf("group = %+v", g)
	}
}

func TestBuildRouteGroupFromArgs_SubAppNotIdent(t *testing.T) {
	call, src := firstCallExpr(t, `app.route("/x", "notident");`+"\n")
	if g := buildRouteGroupFromArgs(call, src, "app"); g != nil {
		t.Fatalf("non-ident subapp should be nil, got %+v", g)
	}
}
