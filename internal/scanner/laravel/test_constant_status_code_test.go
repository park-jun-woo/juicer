//ff:func feature=scan type=test topic=laravel control=sequence
//ff:what constantStatusCode Response::HTTP_* 상수 → 숫자 코드 매핑 테스트
package laravel

import "testing"

func TestConstantStatusCode(t *testing.T) {
	fi := mustParsePHP(t, `<?php abort(Response::HTTP_NOT_FOUND);`)
	args := findAllByType(fi.root, "argument")
	if len(args) == 0 {
		t.Fatal("no argument")
	}
	if got := constantStatusCode(args[0], fi.src); got != "404" {
		t.Errorf("got %q, want 404", got)
	}
	// non-constant argument
	fi2 := mustParsePHP(t, `<?php abort(500);`)
	a2 := findAllByType(fi2.root, "argument")[0]
	if got := constantStatusCode(a2, fi2.src); got != "" {
		t.Errorf("numeric literal: got %q", got)
	}
}
