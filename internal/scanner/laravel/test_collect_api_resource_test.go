//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what Route::apiResource() 기본 CRUD 자동 생성 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAPIResource_Basic(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::apiResource('posts', PostController::class);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := collectAPIResource(*fi, "api", nil)
	if len(routes) != 6 {
		for _, r := range routes {
			t.Logf("  %s %s (%s@%s)", r.method, r.path, r.controller, r.action)
		}
		t.Fatalf("expected 6 routes, got %d", len(routes))
	}

	expected := []struct {
		method string
		path   string
		action string
	}{
		{"GET", "/api/posts", "index"},
		{"POST", "/api/posts", "store"},
		{"GET", "/api/posts/{post}", "show"},
		{"PUT", "/api/posts/{post}", "update"},
		{"PATCH", "/api/posts/{post}", "update"},
		{"DELETE", "/api/posts/{post}", "destroy"},
	}
	for i, exp := range expected {
		if routes[i].method != exp.method || routes[i].path != exp.path || routes[i].action != exp.action {
			t.Errorf("route[%d] = %s %s @%s, want %s %s @%s",
				i, routes[i].method, routes[i].path, routes[i].action,
				exp.method, exp.path, exp.action)
		}
	}
}
