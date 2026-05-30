//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestWalkChain_Scoped 테스트
package laravel

import "testing"

func TestWalkChain_Scoped(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	prefix := ""
	var mw []string
	walkChain(call, fi, &prefix, &mw)
	if prefix != "admin" {
		t.Fatalf("prefix: %q", prefix)
	}
}
