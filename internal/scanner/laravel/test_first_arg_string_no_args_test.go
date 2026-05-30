//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFirstArgString_NoArgs 테스트
package laravel

import "testing"

func TestFirstArgString_NoArgs(t *testing.T) {
	fi := mustParsePHP(t, `<?php foo();`)
	args := findAllByType(fi.root, "arguments")[0]
	if _, ok := firstArgString(args, fi.src); ok {
		t.Fatal("expected false")
	}
}
