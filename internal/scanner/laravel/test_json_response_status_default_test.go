//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJSONResponseStatus_Default 테스트
package laravel

import "testing"

func TestJSONResponseStatus_Default(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([]); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	if got := jsonResponseStatus(ret, fi.src); got != "200" {
		t.Fatalf("got %q", got)
	}
}
