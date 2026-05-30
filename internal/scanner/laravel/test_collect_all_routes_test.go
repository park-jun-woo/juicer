//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestCollectAllRoutes 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutes(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::get('/users', [UserController::class, 'index']);
Route::post('/users', [UserController::class, 'store']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	parsed := map[string]*fileInfo{"routes/api.php": fi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	for _, r := range routes {
		if r.path != "/api/users" {
			t.Errorf("expected api prefix, got %q", r.path)
		}
	}
}
