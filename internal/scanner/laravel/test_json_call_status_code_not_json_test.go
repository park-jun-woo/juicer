//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJSONCallStatusCode_NotJson 테스트
package laravel

import "testing"

func TestJSONCallStatusCode_NotJson(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = $obj->other(1, 2);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
