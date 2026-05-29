//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what apiResource 경로형 리소스명({product}/reviews) 파라미터 처리 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAPIResource_PathParamName(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::apiResource('{product}/reviews', ReviewController::class);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := collectAPIResource(*fi, "api", nil)
	if len(routes) != 6 {
		for _, r := range routes {
			t.Logf("  %s %s", r.method, r.path)
		}
		t.Fatalf("expected 6 routes, got %d", len(routes))
	}
	for _, r := range routes {
		if r.action == "index" && r.path != "/api/{product}/reviews" {
			t.Errorf("index path = %q, want %q", r.path, "/api/{product}/reviews")
		}
		if r.action == "show" && r.path != "/api/{product}/reviews/{review}" {
			t.Errorf("show path = %q, want %q", r.path, "/api/{product}/reviews/{review}")
		}
	}
}
