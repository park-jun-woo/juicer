//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractClosureBody_AnonFunc 테스트
package laravel

import "testing"

func TestExtractClosureBody_AnonFunc(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::group([], function () { $x = 1; });`)
	args := findAllByType(fi.root, "arguments")[0]
	if extractClosureBody(args, fi) == nil {
		t.Fatal("expected closure body")
	}
}
