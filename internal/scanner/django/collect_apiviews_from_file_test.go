//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectAPIViewsFromFile — 단일 파일 APIView 수집을 검증
package django

import "testing"

func TestCollectAPIViewsFromFile(t *testing.T) {
	src := `
class PingView(APIView):
    def get(self, request):
        return Response()

class NotAView:
    pass
`
	fi := newTestFileInfo(t, src)
	views := collectAPIViewsFromFile(fi)
	if len(views) != 1 {
		t.Fatalf("expected 1 APIView (non-view class skipped), got %d", len(views))
	}
	if views[0].name != "PingView" {
		t.Errorf("name = %q, want PingView", views[0].name)
	}
}

func TestCollectAPIViewsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "def f():\n    pass\n")
	if v := collectAPIViewsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
