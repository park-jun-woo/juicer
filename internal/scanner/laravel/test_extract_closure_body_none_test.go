//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractClosureBody_None 테스트
package laravel

import "testing"

func TestExtractClosureBody_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group([], $cb);`)
	args := findAllByType(fi.root, "arguments")[0]
	if extractClosureBody(args, fi) != nil {
		t.Fatal("expected nil")
	}
}
