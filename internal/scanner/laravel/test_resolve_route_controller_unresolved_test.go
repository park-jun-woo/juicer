//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveRouteController_Unresolved 테스트
package laravel

import "testing"

func TestResolveRouteController_Unresolved(t *testing.T) {
	ri := routeInfo{controller: "Missing", action: "show"}
	if resolveRouteController(t.TempDir(), ri, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil when controller not found")
	}
}
