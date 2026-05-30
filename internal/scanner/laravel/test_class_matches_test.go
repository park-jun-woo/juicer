//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestClassMatches 테스트
package laravel

import "testing"

func TestClassMatches(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController {}`)
	if !classMatches(&fi, "UserController") {
		t.Fatal("expected match")
	}
	if classMatches(&fi, "Other") {
		t.Fatal("unexpected match")
	}
}
