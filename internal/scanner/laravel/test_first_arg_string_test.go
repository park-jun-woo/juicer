//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFirstArgString 테스트
package laravel

import "testing"

func TestFirstArgString(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin');`)
	args := findAllByType(fi.root, "arguments")[0]
	s, ok := firstArgString(args, fi.src)
	if !ok || s != "admin" {
		t.Fatalf("got %q %v", s, ok)
	}
}
