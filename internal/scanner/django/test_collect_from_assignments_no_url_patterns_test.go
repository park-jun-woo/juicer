//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFromAssignments_NoURLPatterns 테스트
package django

import "testing"

func TestCollectFromAssignments_NoURLPatterns(t *testing.T) {
	fi := newTestFileInfo(t, "config = {}\n")
	if e := collectFromAssignments(fi); len(e) != 0 {
		t.Fatalf("expected no entries, got %d", len(e))
	}
}
