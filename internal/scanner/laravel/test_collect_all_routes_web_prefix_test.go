//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutes_WebPrefix 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAllRoutes_WebPrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/web.php", `<?php
Route::get('/home', [HomeController::class, 'index']);
`)
	fi, _ := parseFile(dir, filepath.Join(dir, "routes/web.php"))
	parsed := map[string]*fileInfo{"routes/web.php": fi}
	routes := collectAllRoutes(parsed)
	if len(routes) != 1 || routes[0].path != "/home" {
		t.Fatalf("web route: %+v", routes)
	}
}
