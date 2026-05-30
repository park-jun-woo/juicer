//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestIsURLPatternsAssignment 테스트
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
