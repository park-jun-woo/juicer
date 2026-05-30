//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractMethodParams 테스트
package laravel

import "testing"

func TestExtractMethodParams(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id, string $name) {} }`)
	fp := findAllByType(fi.root, "formal_parameters")[0]
	params := extractMethodParams(fp, fi.src)
	if len(params) != 2 || params[0].name != "id" || params[0].typeName != "int" {
		t.Fatalf("got %+v", params)
	}
}
