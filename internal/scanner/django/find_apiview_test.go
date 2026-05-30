//ff:func feature=scan type=test control=sequence topic=django
//ff:what findAPIView — 이름으로 APIView 검색 분기를 검증
package django

import "testing"

func TestFindAPIView(t *testing.T) {
	views := []apiviewInfo{{name: "A"}, {name: "B"}}
	if av := findAPIView(views, "B"); av == nil || av.name != "B" {
		t.Fatalf("expected to find B, got %v", av)
	}
	if av := findAPIView(views, "Missing"); av != nil {
		t.Fatalf("expected nil for missing, got %v", av)
	}
}
