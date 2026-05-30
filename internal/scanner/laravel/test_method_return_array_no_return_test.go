//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestMethodReturnArray_NoReturn 테스트
package laravel

import "testing"

func TestMethodReturnArray_NoReturn(t *testing.T) {
	fi := mustParsePHP(t, `<?php class C { public function m() { $x = 1; } }`)
	method := findAllByType(fi.root, "method_declaration")[0]
	if methodReturnArray(method) != nil {
		t.Fatal("expected nil")
	}
}
