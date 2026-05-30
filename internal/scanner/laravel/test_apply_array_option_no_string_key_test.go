//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyArrayOption_NoStringKey 테스트
package laravel

import "testing"

func TestApplyArrayOption_NoStringKey(t *testing.T) {

	fi := mustParsePHP(t, `<?php $x = [42];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	if len(elems) == 0 {
		t.Skip("no element")
	}
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if prefix != "" || len(mw) != 0 {
		t.Fatalf("unexpected change: prefix=%q mw=%v", prefix, mw)
	}
}
