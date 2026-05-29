//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Route::middleware('auth')->get('/user', handler) 체인 단일 라우트 추출 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestExtractChainedRoute_MiddlewareGet(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::middleware('auth:api')->get('/user', [UserController::class, 'show']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := extractRouteGroups(*fi, "api", nil)
	if len(routes) != 1 {
		for _, r := range routes {
			t.Logf("  %s %s mw=%v", r.method, r.path, r.middleware)
		}
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.method != "GET" || r.path != "/api/user" {
		t.Errorf("route = %s %s, want GET /api/user", r.method, r.path)
	}
	if len(r.middleware) != 1 || r.middleware[0] != "auth:api" {
		t.Errorf("middleware = %v, want [auth:api]", r.middleware)
	}
}
