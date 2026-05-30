//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestGroupArrayValue_NoArrow 테스트
package laravel

import "testing"

func TestGroupArrayValue_NoArrow(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['v1'];`)
	elem := findAllByType(fi.root, "array_element_initializer")[0]
	if v := groupArrayValue(elem); v != nil {
		t.Fatalf("expected nil, got %v", v)
	}
}
