//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectURLs — 모듈별 urlpatterns 수집 분기를 검증
package django

import "testing"

func TestCollectURLs(t *testing.T) {
	withURLs := newTestFileInfo(t, "urlpatterns = [path('a/', v1)]\n")
	withURLs.module = "app.urls"
	empty := newTestFileInfo(t, "x = 1\n")
	empty.module = "app.empty"

	byModule := collectURLs([]fileInfo{withURLs, empty})
	if len(byModule) != 1 {
		t.Fatalf("expected 1 module with urls, got %d: %v", len(byModule), byModule)
	}
	if len(byModule["app.urls"]) != 1 {
		t.Errorf("expected 1 entry for app.urls, got %d", len(byModule["app.urls"]))
	}
	if _, ok := byModule["app.empty"]; ok {
		t.Error("empty module should not be present")
	}
}
