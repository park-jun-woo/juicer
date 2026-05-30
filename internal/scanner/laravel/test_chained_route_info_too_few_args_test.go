//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestChainedRouteInfo_TooFewArgs 테스트
package laravel

import "testing"

func TestChainedRouteInfo_TooFewArgs(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::get('/users');`)
	scoped := findAllByType(fi.root, "scoped_call_expression")
	if len(scoped) == 0 {
		t.Skip("no scoped call")
	}
	if r := chainedRouteInfo(scoped[0], fi, "GET", "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
