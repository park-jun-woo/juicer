//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFromAugmentedAssignments_NoList 테스트
package django

import "testing"

func TestCollectFromAugmentedAssignments_NoList(t *testing.T) {
	fi := newTestFileInfo(t, "urlpatterns += other_patterns\n")
	if e := collectFromAugmentedAssignments(fi); len(e) != 0 {
		t.Fatalf("expected 0 (no list literal), got %d", len(e))
	}
}
