//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestNodeText 테스트
package laravel

import "testing"

func TestNodeText(t *testing.T) {
	fi := mustParsePHP(t, `<?php $hello;`)
	vars := findAllByType(fi.root, "variable_name")
	if len(vars) == 0 {
		t.Skip("no var")
	}
	if got := nodeText(vars[0], fi.src); got != "$hello" {
		t.Fatalf("got %q", got)
	}
}
