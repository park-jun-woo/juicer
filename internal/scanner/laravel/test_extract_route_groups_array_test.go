//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what Route::group(['prefix'=>..,'middleware'=>..], fn) 배열 옵션 그룹 추출 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestExtractRouteGroups_ArrayOptions(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::group(['prefix' => 'products', 'middleware' => ['auth:api']], function () {
    Route::apiResource('reviews', ReviewController::class);
});
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := extractRouteGroups(*fi, "api", nil)
	if len(routes) != 6 {
		for _, r := range routes {
			t.Logf("  %s %s mw=%v", r.method, r.path, r.middleware)
		}
		t.Fatalf("expected 6 routes, got %d", len(routes))
	}
	for _, r := range routes {
		if r.action == "index" && r.path != "/api/products/reviews" {
			t.Errorf("index path = %q, want %q", r.path, "/api/products/reviews")
		}
		if r.action == "show" && r.path != "/api/products/reviews/{review}" {
			t.Errorf("show path = %q, want %q", r.path, "/api/products/reviews/{review}")
		}
		if len(r.middleware) != 1 || r.middleware[0] != "auth:api" {
			t.Errorf("middleware = %v, want [auth:api]", r.middleware)
		}
	}
}
