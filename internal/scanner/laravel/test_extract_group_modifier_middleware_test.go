//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractGroupModifier_Middleware 테스트
package laravel

import "testing"

func TestExtractGroupModifier_Middleware(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::middleware(['auth']);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "" || len(mw) != 1 {
		t.Fatalf("got %q %v", p, mw)
	}
}
