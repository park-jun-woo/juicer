//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExpandAPIResource_NotApiResource 테스트
package laravel

import "testing"

func TestExpandAPIResource_NotApiResource(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', [C::class, 'm']);`)
	calls := findAllByType(fi.root, "scoped_call_expression")
	if r := expandAPIResource(calls[0], fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
