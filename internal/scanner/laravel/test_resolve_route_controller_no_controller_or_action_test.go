//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestResolveRouteController_NoControllerOrAction 테스트
package laravel

import "testing"

func TestResolveRouteController_NoControllerOrAction(t *testing.T) {
	if resolveRouteController("/root", routeInfo{}, map[string]*fileInfo{}) != nil {
		t.Fatal("expected nil for empty controller/action")
	}
}
