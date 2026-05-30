//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestGroupArrayValue 테스트
package laravel

import "testing"

func TestGroupArrayValue(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'v1'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	v := groupArrayValue(elem)
	if v == nil || extractStringContent(v, fi.src) != "v1" {
		t.Fatalf("got %v", v)
	}
}
