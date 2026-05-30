//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestJSONCallStatusCode_OneArg 테스트
package laravel

import "testing"

func TestJSONCallStatusCode_OneArg(t *testing.T) {
	fi := mustParsePHP(t, `<?php $x = response()->json([]);`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if got := jsonCallStatusCode(mcs[0], fi.src); got != "" {
		t.Fatalf("expected empty for single arg, got %q", got)
	}
}
