//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveRouteController 테스트
package laravel

import "testing"

func TestResolveRouteController(t *testing.T) {
	fi := mustParsePHP(t, `<?php class UserController { public function show(int $id) { return $id; } }`)
	parsed := map[string]*fileInfo{"x.php": &fi}
	ri := routeInfo{controller: "UserController", action: "show"}
	cm := resolveRouteController("/root", ri, parsed)
	if cm == nil || cm.name != "show" {
		t.Fatalf("got %+v", cm)
	}
}
