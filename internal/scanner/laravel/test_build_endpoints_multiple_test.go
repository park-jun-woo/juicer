//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildEndpoints_Multiple 테스트
package laravel

import "testing"

func TestBuildEndpoints_Multiple(t *testing.T) {
	routes := []routeInfo{
		{method: "GET", path: "/a"},
		{method: "POST", path: "/b"},
	}
	eps := buildEndpoints(t.TempDir(), routes, map[string]*fileInfo{})
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}
