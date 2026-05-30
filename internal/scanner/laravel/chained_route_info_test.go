//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what chainedRouteInfo 테스트
package laravel

import "testing"

func TestChainedRouteInfo_Basic(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/users', [UserController::class, 'index']);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	r := chainedRouteInfo(scoped[0], fi, "GET", "api", []string{"web"})
	if r == nil {
		t.Fatal("nil route")
	}
	if r.method != "GET" || r.path != "/api/users" {
		t.Fatalf("got %+v", r)
	}
	if r.controller != "UserController" || r.action != "index" {
		t.Fatalf("ctrl/action: %+v", r)
	}
	if len(r.middleware) != 1 || r.middleware[0] != "web" {
		t.Fatalf("mw: %v", r.middleware)
	}
}

func TestChainedRouteInfo_TooFewArgs(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/users');`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if r := chainedRouteInfo(scoped[0], fi, "GET", "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}

func TestChainedRouteInfo_EmptyPath(t *testing.T) {
	// empty string literal path -> extractStringContent == "" -> nil
	fi := mustParsePHP(t, `<?php Route::get('', [C::class, 'm']);`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if r := chainedRouteInfo(scoped[0], fi, "GET", "", nil); r != nil {
		t.Fatalf("expected nil for empty path, got %+v", r)
	}
}
