//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectURLsFromFile — 대입/증분대입 urlpatterns 결합 수집을 검증
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

func TestCollectURLsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if e := collectURLsFromFile(fi); len(e) != 0 {
		t.Fatalf("expected none, got %d", len(e))
	}
}
