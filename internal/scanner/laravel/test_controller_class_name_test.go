//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestControllerClassName 테스트
package laravel

import "testing"

func TestControllerClassName(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = [UserController::class, 'index'];`)
	elems := findAllByType(fi.root, "array_creation_expression")
	if len(elems) == 0 {
		t.Fatal("no array")
	}
	if got := controllerClassName(elems[0], fi.src); got != "UserController" {
		t.Fatalf("got %q", got)
	}
}
