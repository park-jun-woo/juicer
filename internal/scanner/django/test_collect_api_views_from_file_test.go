//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectAPIViewsFromFile 테스트
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
	views := collectAPIViewsFromFile(fi, nil)
	if len(views) != 1 {
		t.Fatalf("expected 1 APIView (non-view class skipped), got %d", len(views))
	}
	if views[0].name != "PingView" {
		t.Errorf("name = %q, want PingView", views[0].name)
	}
}
