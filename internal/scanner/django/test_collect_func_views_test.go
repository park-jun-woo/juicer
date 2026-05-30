//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFuncViews 테스트
package django

import "testing"

func TestCollectFuncViews(t *testing.T) {
	src := `
@api_view(['GET', 'POST'])
def health(request):
    return Response()
`
	fi := newTestFileInfo(t, src)
	views := collectFuncViews([]fileInfo{fi})
	if len(views) != 1 {
		t.Fatalf("expected 1 func view, got %d", len(views))
	}
	if views[0].name != "health" {
		t.Errorf("name = %q, want health", views[0].name)
	}
}
