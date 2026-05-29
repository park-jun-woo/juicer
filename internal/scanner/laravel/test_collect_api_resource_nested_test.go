//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Route::apiResource() 중첩 리소스 경로 생성 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestCollectAPIResource_Nested(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::apiResource('users.posts', PostController::class);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}
	routes := collectAPIResource(*fi, "api", nil)
	if len(routes) != 5 {
		t.Fatalf("expected 5 routes, got %d", len(routes))
	}
	if routes[0].path != "/api/users/{user}/posts" {
		t.Errorf("route[0].path = %q, want %q", routes[0].path, "/api/users/{user}/posts")
	}
	if routes[2].path != "/api/users/{user}/posts/{post}" {
		t.Errorf("route[2].path = %q, want %q", routes[2].path, "/api/users/{user}/posts/{post}")
	}
}
