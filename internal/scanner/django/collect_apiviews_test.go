//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectAPIViews — 여러 파일에서 APIView 수집을 검증
package django

import "testing"

func newTestFileInfo(t *testing.T, src string) fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{relPath: "app/views.py", module: "app.views", src: []byte(src), root: root}
}

func TestCollectAPIViews(t *testing.T) {
	src := `
class PingView(APIView):
    def get(self, request):
        return Response()
`
	fi := newTestFileInfo(t, src)
	views := collectAPIViews([]fileInfo{fi})
	if len(views) != 1 {
		t.Fatalf("expected 1 APIView, got %d", len(views))
	}
	if views[0].name != "PingView" {
		t.Errorf("name = %q, want PingView", views[0].name)
	}
}

func TestCollectAPIViews_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if views := collectAPIViews([]fileInfo{fi}); len(views) != 0 {
		t.Fatalf("expected no views, got %d", len(views))
	}
}
