//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what nestedGroupRoutes 테스트 헬퍼
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
