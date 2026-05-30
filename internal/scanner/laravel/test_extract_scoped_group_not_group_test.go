//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractScopedGroup_NotGroup 테스트
package laravel

import "testing"

func TestExtractScopedGroup_NotGroup(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractScopedGroup(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
