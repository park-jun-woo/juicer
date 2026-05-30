//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFromAugmentedAssignments_WrongName 테스트
package django

import "testing"

func TestCollectFromAugmentedAssignments_WrongName(t *testing.T) {
	fi := newTestFileInfo(t, "routes += [path('a/', v1)]\n")
	if e := collectFromAugmentedAssignments(fi); len(e) != 0 {
		t.Fatalf("expected 0 (wrong target name), got %d", len(e))
	}
}
