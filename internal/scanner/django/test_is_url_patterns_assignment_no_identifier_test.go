//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestIsURLPatternsAssignment_NoIdentifier 테스트
package django

import "testing"

func TestIsURLPatternsAssignment_NoIdentifier(t *testing.T) {

	fi := newTestFileInfo(t, "x = 1\n")
	if isURLPatternsAssignment(fi.root, fi.src) {
		t.Error("expected false for node without identifier")
	}
}
