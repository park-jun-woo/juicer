//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestGroupClosureBody 테스트
package laravel

import "testing"

func TestGroupClosureBody(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('x')->group(function () { $a = 1; });`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	if groupClosureBody(mcs[0], fi) == nil {
		t.Fatal("expected closure body")
	}
}
