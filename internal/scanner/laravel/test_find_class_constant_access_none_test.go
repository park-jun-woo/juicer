//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindClassConstantAccess_None 테스트
package laravel

import "testing"

func TestFindClassConstantAccess_None(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'plain';`)
	str := findAllByType(fi.root, "string")[0]
	if findClassConstantAccess(str) != nil {
		t.Fatal("expected nil")
	}
}
