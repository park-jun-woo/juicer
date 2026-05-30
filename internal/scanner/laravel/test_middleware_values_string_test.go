//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestMiddlewareValues_String 테스트
package laravel

import "testing"

func TestMiddlewareValues_String(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = 'auth';`)
	str := findAllByType(fi.root, "string")[0]
	if got := middlewareValues(str, fi); len(got) != 1 || got[0] != "auth" {
		t.Fatalf("got %v", got)
	}
}
