//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestChildrenOfType 테스트
package laravel

import "testing"

func TestChildrenOfType(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a' => 1, 'b' => 2];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	if len(arrs) == 0 {
		t.Fatal("no array")
	}
	elems := childrenOfType(arrs[0], "array_element_initializer")
	if len(elems) != 2 {
		t.Fatalf("expected 2, got %d", len(elems))
	}
}
