//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyArrayOption_Middleware 테스트
package laravel

import "testing"

func TestApplyArrayOption_Middleware(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['middleware' => 'auth'];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if len(mw) == 0 {
		t.Fatalf("middleware not applied: %v", mw)
	}
}
