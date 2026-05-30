//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJSONResponseStatus 테스트
package laravel

import "testing"

func TestJSONResponseStatus(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([], 201); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if got := jsonResponseStatus(ret, fi.src); got != "201" {
		t.Fatalf("got %q", got)
	}
}
