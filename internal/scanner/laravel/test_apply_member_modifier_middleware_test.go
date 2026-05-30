//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyMemberModifier_Middleware 테스트
package laravel

import "testing"

func TestApplyMemberModifier_Middleware(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x')->middleware('auth');`)
	calls := findAllByType(fi.root, "member_call_expression")
	if len(calls) == 0 {
		t.Skip("no member call")
	}
	prefix := ""
	var mw []string
	applyMemberModifier(calls[0], fi, &prefix, &mw)
	if len(mw) == 0 {
		t.Fatalf("middleware not applied: %v", mw)
	}
}
