//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectAPIViews 테스트
package django

import "testing"

func TestCollectAPIViews(t *testing.T) {
	src := `
class PingView(APIView):
    def get(self, request):
        return Response()
`
	fi := newTestFileInfo(t, src)
	views := collectAPIViews([]fileInfo{fi}, nil)
	if len(views) != 1 {
		t.Fatalf("expected 1 APIView, got %d", len(views))
	}
	if views[0].name != "PingView" {
		t.Errorf("name = %q, want PingView", views[0].name)
	}
}
