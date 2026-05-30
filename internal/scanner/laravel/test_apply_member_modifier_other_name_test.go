//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyMemberModifier_OtherName 테스트
package laravel

import "testing"

func TestApplyMemberModifier_OtherName(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x')->name('foo');`)
	calls := findAllByType(fi.root, "member_call_expression")
	if len(calls) == 0 {
		t.Skip("no member call")
	}
	prefix := ""
	var mw []string
	applyMemberModifier(calls[0], fi, &prefix, &mw)
	if prefix != "" || len(mw) != 0 {
		t.Fatalf("unexpected change: prefix=%q mw=%v", prefix, mw)
	}
}
