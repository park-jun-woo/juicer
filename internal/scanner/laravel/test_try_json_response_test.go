//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestTryJSONResponse 테스트
package laravel

import "testing"

func TestTryJSONResponse(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { return response()->json([], 201); } }`)
	ret := findAllByType(fi.root, "return_statement")[0]
	resp := tryJSONResponse(ret, fi.src, nodeText(ret, fi.src))
	if resp == nil || resp.Status != "201" || resp.Kind != "json" {
		t.Fatalf("got %+v", resp)
	}
}
