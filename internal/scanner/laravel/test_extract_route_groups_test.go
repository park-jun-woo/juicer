//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what Route::prefix()->group() / Route::middleware()->group() 그룹 라우트 추출 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestExtractRouteGroups_Prefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::prefix('v1')->group(function () {
    Route::get('/health', function () { return response()->json(['status' => 'ok']); });
});
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := extractRouteGroups(*fi, "api", nil)
	if len(routes) != 1 {
		for _, r := range routes {
			t.Logf("  %s %s", r.method, r.path)
		}
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/api/v1/health" {
		t.Errorf("path = %q, want %q", routes[0].path, "/api/v1/health")
	}
}

func TestExtractRouteGroups_Middleware(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::middleware(['auth:sanctum'])->group(function () {
    Route::get('/me', [ProfileController::class, 'show']);
});
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := extractRouteGroups(*fi, "api", nil)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if routes[0].path != "/api/me" {
		t.Errorf("path = %q, want %q", routes[0].path, "/api/me")
	}
	if len(routes[0].middleware) != 1 || routes[0].middleware[0] != "auth:sanctum" {
		t.Errorf("middleware = %v, want [auth:sanctum]", routes[0].middleware)
	}
}
