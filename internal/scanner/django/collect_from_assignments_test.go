//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectFromAssignments — urlpatterns 대입에서 path 수집 분기를 검증
package django

import "testing"

func TestCollectFromAssignments(t *testing.T) {
	src := `
other = 5
urlpatterns = [path('a/', v1), path('b/', v2)]
`
	fi := newTestFileInfo(t, src)
	entries := collectFromAssignments(fi)
	if len(entries) != 2 {
		t.Fatalf("expected 2 url entries, got %d", len(entries))
	}
}

func TestCollectFromAssignments_NoURLPatterns(t *testing.T) {
	fi := newTestFileInfo(t, "config = {}\n")
	if e := collectFromAssignments(fi); len(e) != 0 {
		t.Fatalf("expected no entries, got %d", len(e))
	}
}
