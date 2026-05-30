//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractChainedGroupIfPresent_NoInner 테스트
package laravel

import "testing"

func TestExtractChainedGroupIfPresent_NoInner(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', [C::class, 'm']);`)

	scoped := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractChainedGroupIfPresent(scoped, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
