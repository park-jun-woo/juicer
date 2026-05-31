//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectAPIViews_None 테스트
package django

import "testing"

func TestCollectAPIViews_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if views := collectAPIViews([]fileInfo{fi}, nil); len(views) != 0 {
		t.Fatalf("expected no views, got %d", len(views))
	}
}
