//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestMiddlewareValues_Array 테스트
package laravel

import "testing"

func TestMiddlewareValues_Array(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['auth', 'throttle'];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	if got := middlewareValues(arr, fi); len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}
