//ff:func feature=scan type=test control=sequence topic=django
//ff:what include 대상 모듈을 접미사 매칭으로 해석하는지 검증한다
package django

import "testing"

func TestResolveIncludeModule(t *testing.T) {
	byModule := map[string][]urlEntry{
		"config.urls": nil,
		"blog.urls":   nil,
	}
	if got, ok := resolveIncludeModule("blog.urls", byModule); !ok || got != "blog.urls" {
		t.Errorf("exact match failed: got %q ok=%v", got, ok)
	}
	// suffix match: "urls" should not match "blog.urls" unless dotted suffix matches "urls"
	if got, ok := resolveIncludeModule("urls", byModule); !ok || (got != "config.urls" && got != "blog.urls") {
		t.Errorf("suffix match failed: got %q ok=%v", got, ok)
	}
	if _, ok := resolveIncludeModule("nope.urls", byModule); ok {
		t.Errorf("expected no match for nope.urls")
	}
	if _, ok := resolveIncludeModule("", byModule); ok {
		t.Errorf("expected no match for empty target")
	}
}
