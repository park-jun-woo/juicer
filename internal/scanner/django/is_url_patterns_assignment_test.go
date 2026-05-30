//ff:func feature=scan type=test control=sequence topic=django
//ff:what isURLPatternsAssignment — urlpatterns 대입 판별 분기를 검증
package django

import "testing"

func TestIsURLPatternsAssignment(t *testing.T) {
	yes := newTestFileInfo(t, "urlpatterns = [path('a/', v)]\n")
	for _, a := range findAllByType(yes.root, "assignment") {
		if !isURLPatternsAssignment(a, yes.src) {
			t.Error("expected true for urlpatterns assignment")
		}
	}

	no := newTestFileInfo(t, "routes = [1]\n")
	for _, a := range findAllByType(no.root, "assignment") {
		if isURLPatternsAssignment(a, no.src) {
			t.Error("expected false for non-urlpatterns assignment")
		}
	}
}

func TestIsURLPatternsAssignment_NoIdentifier(t *testing.T) {
	// A node without an identifier child -> false. Pass the module root.
	fi := newTestFileInfo(t, "x = 1\n")
	if isURLPatternsAssignment(fi.root, fi.src) {
		t.Error("expected false for node without identifier")
	}
}
