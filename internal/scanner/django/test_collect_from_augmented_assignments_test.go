//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFromAugmentedAssignments 테스트
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
