//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractClassReference 테스트
package laravel

import "testing"

func TestExtractClassReference(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = PostController::class;`)
	exprs := findAllByType(fi.root, "class_constant_access_expression")
	if len(exprs) == 0 {
		t.Fatal("no class const access")
	}
	if got := extractClassReference(exprs[0], fi.src); got != "PostController" {
		t.Fatalf("got %q", got)
	}
}
