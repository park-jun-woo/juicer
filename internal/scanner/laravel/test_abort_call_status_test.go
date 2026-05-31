//ff:func feature=scan type=test topic=laravel control=sequence
//ff:what abortCallStatus abort()/abort_if() 호출의 status 코드 추출 테스트
package laravel

import "testing"

func TestAbortCallStatus(t *testing.T) {
	// abort(404) -> 404 (arg index 0)
	fi := mustParsePHP(t, `<?php abort(404);`)
	call := findAllByType(fi.root, "function_call_expression")[0]
	if got := abortCallStatus(call, fi.src); got != "404" {
		t.Errorf("abort: got %q", got)
	}
	// abort_if(cond, 403) -> 403 (arg index 1)
	fi2 := mustParsePHP(t, `<?php abort_if($x, 403);`)
	c2 := findAllByType(fi2.root, "function_call_expression")[0]
	if got := abortCallStatus(c2, fi2.src); got != "403" {
		t.Errorf("abort_if: got %q", got)
	}
	// non-abort function
	fi3 := mustParsePHP(t, `<?php other(404);`)
	c3 := findAllByType(fi3.root, "function_call_expression")[0]
	if got := abortCallStatus(c3, fi3.src); got != "" {
		t.Errorf("non-abort: got %q", got)
	}
}
