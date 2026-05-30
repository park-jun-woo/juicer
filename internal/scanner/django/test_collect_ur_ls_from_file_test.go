//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectURLsFromFile 테스트
package django

import "testing"

func TestCollectURLsFromFile(t *testing.T) {
	src := `
urlpatterns = [path('a/', v1)]
urlpatterns += [path('b/', v2)]
`
	fi := newTestFileInfo(t, src)
	entries := collectURLsFromFile(fi)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries (assignment + augmented), got %d", len(entries))
	}
}
