//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractControllerAction 테스트
package laravel

import "testing"

func TestExtractControllerAction(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = [UserController::class, 'show'];`)
	arrs := findAllByType(fi.root, "array_creation_expression")
	ctrl, action := extractControllerAction(arrs[0], fi.src)
	if ctrl != "UserController" || action != "show" {
		t.Fatalf("got %q %q", ctrl, action)
	}
}
