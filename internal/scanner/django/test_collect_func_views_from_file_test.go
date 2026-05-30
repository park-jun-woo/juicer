//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFuncViewsFromFile 테스트
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
