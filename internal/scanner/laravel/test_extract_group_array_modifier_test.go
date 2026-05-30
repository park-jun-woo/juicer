//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractGroupArrayModifier 테스트
package laravel

import "testing"

func TestExtractGroupArrayModifier(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['prefix' => 'admin', 'middleware' => ['auth']];`)
	arr := findAllByType(fi.root, "array_creation_expression")[0]
	prefix, mw := extractGroupArrayModifier(arr, fi)
	if prefix == "" || len(mw) != 1 {
		t.Fatalf("got %q %v", prefix, mw)
	}
}
