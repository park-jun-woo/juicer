//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectFromAugmentedAssignments — urlpatterns += [...] 수집 분기를 검증
package django

import "testing"

func TestCollectFromAugmentedAssignments(t *testing.T) {
	src := `
urlpatterns += [path('a/', v1)]
`
	fi := newTestFileInfo(t, src)
	entries := collectFromAugmentedAssignments(fi)
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
}

func TestCollectFromAugmentedAssignments_WrongName(t *testing.T) {
	fi := newTestFileInfo(t, "routes += [path('a/', v1)]\n")
	if e := collectFromAugmentedAssignments(fi); len(e) != 0 {
		t.Fatalf("expected 0 (wrong target name), got %d", len(e))
	}
}

func TestCollectFromAugmentedAssignments_NoList(t *testing.T) {
	fi := newTestFileInfo(t, "urlpatterns += other_patterns\n")
	if e := collectFromAugmentedAssignments(fi); len(e) != 0 {
		t.Fatalf("expected 0 (no list literal), got %d", len(e))
	}
}
