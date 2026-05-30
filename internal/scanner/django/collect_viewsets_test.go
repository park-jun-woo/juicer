//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectViewSets — ViewSet 클래스 수집을 검증
package django

import "testing"

func TestCollectViewSets(t *testing.T) {
	src := `
class UserViewSet(ModelViewSet):
    serializer_class = UserSerializer
`
	fi := newTestFileInfo(t, src)
	vs := collectViewSets([]fileInfo{fi})
	if len(vs) != 1 {
		t.Fatalf("expected 1 viewset, got %d", len(vs))
	}
	if vs[0].name != "UserViewSet" {
		t.Errorf("name = %q, want UserViewSet", vs[0].name)
	}
}

func TestCollectViewSets_None(t *testing.T) {
	fi := newTestFileInfo(t, "class Plain:\n    pass\n")
	if v := collectViewSets([]fileInfo{fi}); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
