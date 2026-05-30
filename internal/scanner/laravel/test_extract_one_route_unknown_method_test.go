//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_UnknownMethod 테스트
package laravel

import "testing"

func TestExtractOneRoute_UnknownMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::macro('/x', $cb);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractOneRoute(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
