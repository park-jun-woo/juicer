//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractMiddlewareArgs_Single 테스트
package laravel

import "testing"

func TestExtractMiddlewareArgs_Single(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware('auth');`)
	args := findAllByType(fi.root, "arguments")[0]
	mw := extractMiddlewareArgs(args, fi)
	if len(mw) != 1 || mw[0] != "auth" {
		t.Fatalf("got %v", mw)
	}
}
