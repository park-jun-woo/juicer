//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what 그룹/체인 헬퍼 공용 픽스처 (round5)
package laravel

import (
	"path/filepath"
	"testing"
)

// nestedGroupRoutes parses a routes file with a chained
// prefix()->middleware()->group(fn) call containing a nested
// prefix()->group(fn), exercising the recursive group walk helpers.
func nestedGroupRoutes(t *testing.T) []routeInfo {
	t.Helper()
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::prefix('v1')->middleware(['auth'])->group(function () {
    Route::get('/health', function () { return response()->json([]); });
    Route::prefix('admin')->group(function () {
        Route::get('/stats', function () { return response()->json([]); });
    });
});
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	return extractRouteGroups(*fi, "api", nil)
}

func TestExtractChainedRoute_Round5(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::middleware('auth')->get('/profile', [UserController::class, 'profile']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	// find the outer member_call_expression (the ->get(...) chain)
	mcs := findAllByType(fi.root, "member_call_expression")
	var got *routeInfo
	for _, mc := range mcs {
		if r := extractChainedRoute(mc, *fi, "api", nil); r != nil {
			got = r
			break
		}
	}
	if got == nil {
		t.Fatal("expected a chained route")
	}
	if got.method != "GET" {
		t.Errorf("method: %q", got.method)
	}
	if len(got.middleware) == 0 || got.middleware[0] != "auth" {
		t.Errorf("middleware: %v", got.middleware)
	}
}

func findRoute(routes []routeInfo, path string) (routeInfo, bool) {
	for _, r := range routes {
		if r.path == path {
			return r, true
		}
	}
	return routeInfo{}, false
}
