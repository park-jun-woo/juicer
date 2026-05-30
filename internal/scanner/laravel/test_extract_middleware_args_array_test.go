//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractMiddlewareArgs_Array 테스트
package laravel

import "testing"

func TestExtractMiddlewareArgs_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware(['auth', 'throttle']);`)
	args := findAllByType(fi.root, "arguments")[0]
	mw := extractMiddlewareArgs(args, fi)
	if len(mw) != 2 {
		t.Fatalf("got %v", mw)
	}
}
