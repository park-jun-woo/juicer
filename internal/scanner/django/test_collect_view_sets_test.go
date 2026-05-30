//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectViewSets 테스트
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
