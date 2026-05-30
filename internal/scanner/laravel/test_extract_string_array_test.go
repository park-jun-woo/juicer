//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractStringArray 테스트
package laravel

import "testing"

func TestExtractStringArray(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a', 'b', 'c'];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	got := extractStringArray(arr, fi.src)
	if len(got) != 3 || got[0] != "a" {
		t.Fatalf("got %v", got)
	}
}
