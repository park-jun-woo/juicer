//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectViewSetsFromFile — 단일 파일 ViewSet 수집을 검증
package django

import "testing"

func TestCollectViewSetsFromFile(t *testing.T) {
	src := `
class UserViewSet(ModelViewSet):
    pass

class Plain:
    pass
`
	fi := newTestFileInfo(t, src)
	vs := collectViewSetsFromFile(fi)
	if len(vs) != 1 {
		t.Fatalf("expected 1 viewset (Plain skipped), got %d", len(vs))
	}
	if vs[0].name != "UserViewSet" {
		t.Errorf("name = %q, want UserViewSet", vs[0].name)
	}
}

func TestCollectViewSetsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if v := collectViewSetsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
