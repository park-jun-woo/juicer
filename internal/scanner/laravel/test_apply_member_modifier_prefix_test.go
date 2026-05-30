//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyMemberModifier_Prefix 테스트
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

	applyMemberModifier(calls[0], fi, &prefix, &mw)
	if prefix != "/v1" && prefix != "v1" {
		t.Fatalf("prefix got %q", prefix)
	}
}
