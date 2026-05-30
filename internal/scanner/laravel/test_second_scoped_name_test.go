//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestSecondScopedName 테스트
package laravel

import "testing"

func TestSecondScopedName(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if got := secondScopedName(call, fi.src); got != "get" {
		t.Fatalf("got %q", got)
	}
}
