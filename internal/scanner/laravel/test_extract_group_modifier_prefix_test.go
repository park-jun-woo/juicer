//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractGroupModifier_Prefix 테스트
package laravel

import "testing"

func TestExtractGroupModifier_Prefix(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "admin" || mw != nil {
		t.Fatalf("got %q %v", p, mw)
	}
}
