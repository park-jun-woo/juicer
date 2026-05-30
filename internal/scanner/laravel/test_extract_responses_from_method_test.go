//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractResponsesFromMethod 테스트
package laravel

import "testing"

func TestExtractResponsesFromMethod(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() {
		return response()->json(['ok' => true], 201);
	} }`)
	cm := extractControllerMethod(&fi, "m")
	resps := extractResponsesFromMethod(t.TempDir(), cm, map[string]*fileInfo{})
	if len(resps) == 0 || resps[0].Status != "201" {
		t.Fatalf("got %+v", resps)
	}
}
