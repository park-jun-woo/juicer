//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Route::get/post 개별 라우트 수집 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectRoutes_Basic(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
Route::post('/users', [UserController::class, 'store']);
Route::get('/users/{user}', [UserController::class, 'show']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := collectRoutes(*fi, "api", nil)
	if len(routes) != 3 {
		t.Fatalf("expected 3 routes, got %d", len(routes))
	}
	if routes[0].method != "GET" || routes[0].path != "/api/users" {
		t.Errorf("route[0] = %s %s, want GET /api/users", routes[0].method, routes[0].path)
	}
	if routes[1].method != "POST" || routes[1].path != "/api/users" {
		t.Errorf("route[1] = %s %s, want POST /api/users", routes[1].method, routes[1].path)
	}
	if routes[2].method != "GET" || routes[2].path != "/api/users/{user}" {
		t.Errorf("route[2] = %s %s, want GET /api/users/{user}", routes[2].method, routes[2].path)
	}
	if routes[0].controller != "UserController" || routes[0].action != "index" {
		t.Errorf("route[0] handler = %s@%s, want UserController@index", routes[0].controller, routes[0].action)
	}
}
