//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestControllerClassName_NoClassAccess 테스트
package laravel

import "testing"

func TestControllerClassName_NoClassAccess(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = ['a', 'b'];`)
	elems := findAllByType(fi.root, "array_creation_expression")
	if got := controllerClassName(elems[0], fi.src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
