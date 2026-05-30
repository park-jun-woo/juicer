//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractScopedGroup 테스트
package laravel

import "testing"

func TestExtractScopedGroup(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group(['prefix' => 'admin'], function () {
		Route::get('/users', [C::class, 'index']);
	});`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	routes := extractScopedGroup(call, fi, "", nil)
	if len(routes) == 0 {
		t.Fatalf("expected routes, got %+v", routes)
	}
	if routes[0].path != "/admin/users" {
		t.Fatalf("path: %q", routes[0].path)
	}
}
