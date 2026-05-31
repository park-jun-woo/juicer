//ff:func feature=scan type=test topic=laravel control=sequence
//ff:what resolveStatusArg 상수/리터럴 status 인자 해석 테스트
package laravel

import "testing"

func TestResolveStatusArg(t *testing.T) {
	// constant
	fi := mustParsePHP(t, `<?php abort(Response::HTTP_FORBIDDEN);`)
	a := findAllByType(fi.root, "argument")[0]
	if got := resolveStatusArg(a, fi.src); got != "403" {
		t.Errorf("constant: got %q", got)
	}
	// numeric literal
	fi2 := mustParsePHP(t, `<?php abort(418);`)
	a2 := findAllByType(fi2.root, "argument")[0]
	if got := resolveStatusArg(a2, fi2.src); got != "418" {
		t.Errorf("literal: got %q", got)
	}
}
