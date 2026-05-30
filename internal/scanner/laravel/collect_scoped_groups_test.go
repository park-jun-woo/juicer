//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCollectScopedGroups_Round5 (round5) — 중첩/체인 그룹 추출로 헬퍼 검증
package laravel

import "testing"

func TestCollectScopedGroups_Round5(t *testing.T) {
	routes := nestedGroupRoutes(t)
	if _, ok := findRoute(routes, "/api/v1/health"); !ok {
		t.Fatalf("expected /api/v1/health route, got %d routes", len(routes))
	}
	if _, ok := findRoute(routes, "/api/v1/admin/stats"); !ok {
		t.Fatalf("expected nested /api/v1/admin/stats route, got %d routes", len(routes))
	}
	for _, r := range routes {
		if len(r.middleware) == 0 || r.middleware[0] != "auth" {
			t.Errorf("route %s should carry auth middleware, got %v", r.path, r.middleware)
		}
	}
}
