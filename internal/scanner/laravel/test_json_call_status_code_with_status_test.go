//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJSONCallStatusCode_WithStatus 테스트
package laravel

import "testing"

func TestJSONCallStatusCode_WithStatus(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = response()->json([], 201);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "201" {
		t.Fatalf("got %q", got)
	}
}
