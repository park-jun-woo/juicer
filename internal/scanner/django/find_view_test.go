//ff:func feature=scan type=test control=sequence topic=django
//ff:what findViewSet — 이름으로 ViewSet 검색 분기를 검증
package django

import "testing"

func TestFindViewSet(t *testing.T) {
	vs := []viewsetInfo{{name: "A"}, {name: "B"}}
	if got := findViewSet(vs, "A"); got == nil || got.name != "A" {
		t.Fatalf("expected to find A, got %v", got)
	}
	if got := findViewSet(vs, "Missing"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
