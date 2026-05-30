//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestExtractChainedRoute_Round5 테스트
package laravel

import (
	"path/filepath"
	"testing"
)

func TestExtractChainedRoute_Round5(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/api.php", `<?php
Route::middleware('auth')->get('/profile', [UserController::class, 'profile']);
`)
	fi, err := parseFile(dir, filepath.Join(dir, "routes/api.php"))
	if err != nil {
		t.Fatal(err)
	}

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
