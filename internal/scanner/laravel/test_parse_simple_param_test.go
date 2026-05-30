//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestParseSimpleParam 테스트
package laravel

import "testing"

func TestParseSimpleParam(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m(int $id) {} }`)
	sp := findAllByType(fi.root, "simple_parameter")[0]
	mp := parseSimpleParam(sp, fi.src)
	if mp.name != "id" || mp.typeName != "int" {
		t.Fatalf("got %+v", mp)
	}
}
