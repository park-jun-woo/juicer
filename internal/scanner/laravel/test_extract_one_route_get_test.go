//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_Get 테스트
package laravel

import "testing"

func TestExtractOneRoute_Get(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/users', [UserController::class, 'index']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	r := extractOneRoute(call, fi, "api", nil)
	if r == nil || r.method != "GET" || r.path != "/api/users" {
		t.Fatalf("got %+v", r)
	}
}
