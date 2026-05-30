//ff:func feature=scan type=test control=sequence topic=django
//ff:what collectIncludedModules — include 참조 모듈 수집 분기를 검증
package django

import "testing"

func TestCollectIncludedModules(t *testing.T) {
	byModule := map[string][]urlEntry{
		"app.urls": {},
	}
	entries := []urlEntry{
		{pattern: "api/", isInclude: true, includeModule: "app.urls"}, // resolves
		{pattern: "x/", isInclude: false},                             // skipped (not include)
		{pattern: "y/", isInclude: true, includeModule: "missing.urls"}, // unresolved
	}
	included := map[string]bool{}
	collectIncludedModules(entries, byModule, included)

	if !included["app.urls"] {
		t.Error("expected app.urls to be marked included")
	}
	if len(included) != 1 {
		t.Errorf("expected exactly 1 included module, got %v", included)
	}
}
