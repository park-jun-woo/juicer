//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractScopedGroup_NotRouteScope 테스트
package laravel

import "testing"

func TestExtractScopedGroup_NotRouteScope(t *testing.T) {
	fi := mustParsePHP(t, `<?php Foo::group(['prefix'=>'x'], fn() => null);`)
	call := findAllByType(fi.root, "scoped_call_expression")[0]
	if r := extractScopedGroup(call, fi, "", nil); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
