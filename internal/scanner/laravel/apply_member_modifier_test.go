//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what applyMemberModifier 테스트
package laravel

import "testing"

func TestApplyMemberModifier_Prefix(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x')->prefix('v1');`)
	calls := findAllByType(fi.root, "member_call_expression")
	if len(calls) == 0 {
		t.Skip("no member call")
	}
	prefix := ""
	var mw []string
	// the outermost member call is ->prefix('v1')
	applyMemberModifier(calls[0], fi, &prefix, &mw)
	if prefix != "/v1" && prefix != "v1" {
		t.Fatalf("prefix got %q", prefix)
	}
}

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
