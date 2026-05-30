//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractOneRoute_NotRouteScope 테스트
package laravel

import "testing"

func TestExtractOneRoute_NotRouteScope(t *testing.T) {
	fi := mustParsePHP(t, `<?php Foo::get('/x', $h);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractOneRoute(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
