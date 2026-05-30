//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractGroupModifier_Other 테스트
package laravel

import "testing"

func TestExtractGroupModifier_Other(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::name('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	p, mw := extractGroupModifier(call, fi)
	if p != "" || mw != nil {
		t.Fatalf("got %q %v", p, mw)
	}
}
