//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what collectAPIViews — 커스텀 BaseAPIView 경유 본문 메서드 전개를 검증
package django

import "testing"

func TestCollectAPIViewsTransitive(t *testing.T) {
	base := newTestFileInfo(t, `
class BaseAPIView(APIView):
    pass
`)
	view := newTestFileInfo(t, `
class WorkspaceEndpoint(BaseAPIView):
    def get(self, request):
        return None

    def post(self, request):
        return None
`)
	files := []fileInfo{base, view}
	idx := buildClassIndex(files)
	views := collectAPIViews(files, idx)

	var av *apiviewInfo
	for i := range views {
		if views[i].name == "WorkspaceEndpoint" {
			av = &views[i]
		}
	}
	if av == nil {
		t.Fatalf("WorkspaceEndpoint not recognized transitively via BaseAPIView; got %+v", views)
	}
	if len(av.methods) != 2 {
		t.Errorf("expected GET+POST methods expanded, got %v", av.methods)
	}
}
