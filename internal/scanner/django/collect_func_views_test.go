//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectFuncViews — @api_view 함수 뷰 수집을 검증
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

func TestCollectFuncViews_None(t *testing.T) {
	fi := newTestFileInfo(t, "def plain(request):\n    return None\n")
	if v := collectFuncViews([]fileInfo{fi}); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
