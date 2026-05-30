//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFuncViews_None 테스트
package django

import "testing"

func TestCollectFuncViews_None(t *testing.T) {
	fi := newTestFileInfo(t, "def plain(request):\n    return None\n")
	if v := collectFuncViews([]fileInfo{fi}); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
