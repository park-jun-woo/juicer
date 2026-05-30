//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectFuncViewsFromFile — 단일 파일 @api_view 함수 뷰 수집을 검증
package django

import "testing"

func TestCollectFuncViewsFromFile(t *testing.T) {
	src := `
@api_view(['GET'])
def health(request):
    return Response()

def plain(request):
    return None
`
	fi := newTestFileInfo(t, src)
	views := collectFuncViewsFromFile(fi)
	if len(views) != 1 {
		t.Fatalf("expected 1 func view (plain skipped), got %d", len(views))
	}
	if views[0].name != "health" {
		t.Errorf("name = %q, want health", views[0].name)
	}
}

func TestCollectFuncViewsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if v := collectFuncViewsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
