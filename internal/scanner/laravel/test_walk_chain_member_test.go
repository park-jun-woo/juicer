//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestWalkChain_Member 테스트
package laravel

import "testing"

func TestWalkChain_Member(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin')->middleware('auth');`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	prefix := ""
	var mw []string
	walkChain(mcs[0], fi, &prefix, &mw)
	if prefix != "admin" || len(mw) == 0 {
		t.Fatalf("prefix=%q mw=%v", prefix, mw)
	}
}
