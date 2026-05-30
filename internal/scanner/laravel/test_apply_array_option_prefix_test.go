//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestApplyArrayOption_Prefix 테스트
package laravel

import "testing"

func TestApplyArrayOption_Prefix(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'v1'];`)
	elems := findAllByType(fi.root, "array_element_initializer")
	if len(elems) == 0 {
		t.Fatal("no array element")
	}
	prefix := ""
	var mw []string
	applyArrayOption(elems[0], fi, &prefix, &mw)
	if prefix != "/v1" && prefix != "v1" {
		t.Fatalf("prefix got %q", prefix)
	}
}
