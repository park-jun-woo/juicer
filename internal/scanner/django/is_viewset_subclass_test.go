//ff:func feature=scan type=test control=sequence topic=django
//ff:what isViewSetSubclass — ViewSet 부모 판별 분기를 검증
package django

import "testing"

func TestIsViewSetSubclass(t *testing.T) {
	if !isViewSetSubclass([]string{"ModelViewSet"}) {
		t.Error("expected true for ModelViewSet")
	}
	if isViewSetSubclass([]string{"object"}) {
		t.Error("expected false for non-viewset parent")
	}
	if isViewSetSubclass(nil) {
		t.Error("expected false for no parents")
	}
}
