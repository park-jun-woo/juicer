//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractMiddlewareArgs_None 테스트
package laravel

import "testing"

func TestExtractMiddlewareArgs_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware();`)
	args := findAllByType(fi.root, "arguments")[0]
	if mw := extractMiddlewareArgs(args, fi); mw != nil {
		t.Fatalf("expected nil, got %v", mw)
	}
}
