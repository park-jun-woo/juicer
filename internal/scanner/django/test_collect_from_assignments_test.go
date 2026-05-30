//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFromAssignments 테스트
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
