//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what apiResource 선행 슬래시 리소스명 정규화 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAPIResource_LeadingSlash(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::apiResource('/products', ProductController::class);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := collectAPIResource(*fi, "api", nil)
	if len(routes) != 6 {
		t.Fatalf("expected 6 routes, got %d", len(routes))
	}
	want := map[string]string{
		"index": "/api/products",
		"show":  "/api/products/{product}",
	}
	for _, r := range routes {
		exp, ok := want[r.action]
		if ok && r.path != exp {
			t.Errorf("%s path = %q, want %q", r.action, r.path, exp)
		}
	}
}
