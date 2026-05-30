//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestLastMemberCallName 테스트
package laravel

import "testing"

func TestLastMemberCallName(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/x')->name('foo');`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := lastMemberCallName(mcs[0], fi.src); got != "name" {
		t.Fatalf("got %q", got)
	}
}
