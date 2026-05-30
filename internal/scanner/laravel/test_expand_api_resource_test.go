//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestExpandAPIResource 테스트
package laravel

import "testing"

func TestExpandAPIResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::apiResource('posts', PostController::class);`)
	calls := findAllByType(fi.root, "scoped_call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	routes := expandAPIResource(calls[0], fi, "api", nil)
	if len(routes) == 0 {
		t.Fatal("expected CRUD routes")
	}
	for _, r := range routes {
		if r.controller != "PostController" {
			t.Errorf("controller %q", r.controller)
		}
	}
}
