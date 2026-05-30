//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestDeduplicateBuilderRoutes_Round5 테스트
package actix

import "testing"

func TestDeduplicateBuilderRoutes_Round5(t *testing.T) {
	routes := []builderRoute{
		{method: "GET", path: "/a", handler: "h"},
		{method: "GET", path: "/a", handler: "h"},
		{method: "POST", path: "/a", handler: "h"},
	}
	got := deduplicateBuilderRoutes(routes)
	if len(got) != 2 {
		t.Fatalf("expected 2 unique, got %d", len(got))
	}
}
