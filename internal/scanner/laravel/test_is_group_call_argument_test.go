//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestIsGroupCallArgument 테스트
package laravel

import "testing"

func TestIsGroupCallArgument(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('x')->group(function () { Route::get('/y', $h); });`)
	closures := findAllByType(fi.root, "anonymous_function_creation_expression")
	if len(closures) == 0 {
		t.Skip("no closure")
	}
	if !isGroupCallArgument(closures[0], fi) {
		t.Fatal("expected true for group closure")
	}
}
