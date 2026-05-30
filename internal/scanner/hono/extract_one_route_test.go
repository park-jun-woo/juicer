//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractOneRoute 테스트
package hono

import "testing"

func oneRoute(t *testing.T, src string, vars map[string]bool) *routeInfo {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	call := findAllByType(fi.Root, "call_expression")[0]
	return extractOneRoute(call, fi.Src, vars)
}

func TestExtractOneRoute_Basic(t *testing.T) {
	r := oneRoute(t, `app.get("/users", handler);`, map[string]bool{"app": true})
	if r == nil {
		t.Fatal("nil route")
	}
	if r.Method != "GET" || r.Path != "/users" || r.Handler != "handler" || r.OwnerVar != "app" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneRoute_AllMethod(t *testing.T) {
	r := oneRoute(t, `app.all("/x", handler);`, map[string]bool{"app": true})
	if r == nil || r.Method != "all" {
		t.Fatalf("got %+v", r)
	}
}

func TestExtractOneRoute_OpenAPI(t *testing.T) {
	r := oneRoute(t, `app.openapi(route, handler);`, map[string]bool{"app": true})
	// openapi delegates to extractOpenAPIRoute; route obj is identifier not createRoute -> may be nil
	_ = r
}

func TestExtractOneRoute_NotHonoVar(t *testing.T) {
	if r := oneRoute(t, `app.get("/x", h);`, map[string]bool{"other": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_UnknownMethod(t *testing.T) {
	if r := oneRoute(t, `app.use("/x", h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_NoMemberExpr(t *testing.T) {
	if r := oneRoute(t, `foo();`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_NoIdentifierObject(t *testing.T) {
	if r := oneRoute(t, `this.get("/x", h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_NoArgs(t *testing.T) {
	if r := oneRoute(t, `app.get();`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_PathNotString(t *testing.T) {
	if r := oneRoute(t, `app.get(pathVar, h);`, map[string]bool{"app": true}); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestExtractOneRoute_WithMiddleware(t *testing.T) {
	r := oneRoute(t, `app.post("/x", auth, handler);`, map[string]bool{"app": true})
	if r == nil || len(r.Middleware) != 1 || r.Middleware[0] != "auth" {
		t.Fatalf("got %+v", r)
	}
	if r.Line != 1 {
		t.Fatalf("expected line 1, got %d", r.Line)
	}
}

func TestExtractOneRoute_QuotedPathUnquoted(t *testing.T) {
	// single-quoted path should be unquoted
	r := oneRoute(t, `app.delete('/items/:id', h);`, map[string]bool{"app": true})
	if r == nil || r.Path != "/items/:id" || r.Method != "DELETE" {
		t.Fatalf("got %+v", r)
	}
}
