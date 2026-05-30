//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectAllRoutes_NoRouteFiles 테스트
package laravel

import "testing"

func TestCollectAllRoutes_NoRouteFiles(t *testing.T) {
	if routes := collectAllRoutes(map[string]*fileInfo{}); len(routes) != 0 {
		t.Fatalf("expected none, got %d", len(routes))
	}
}
