//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestExtractMemberCallRoutes_Group 테스트
package laravel

import "testing"

func TestExtractMemberCallRoutes_Group(t *testing.T) {
	fi := mustParsePHP(t, `<?php Route::prefix('admin')->group(function () {
		Route::get('/users', [C::class, 'index']);
	});`)
	mcs := findAllByType(fi.root, "member_call_expression")
	if len(mcs) == 0 {
		t.Skip("no member call")
	}
	routes := extractMemberCallRoutes(mcs[0], fi, "", nil)
	if len(routes) == 0 {
		t.Fatalf("expected routes, got %+v", routes)
	}
}
