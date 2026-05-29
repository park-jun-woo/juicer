//ff:func feature=scan type=test control=sequence topic=django
//ff:what 단일 urlEntry를 prefix와 결합하고 include면 재귀 전개하는지 검증한다
package django

import "testing"

func TestExpandURLEntry(t *testing.T) {
	byModule := map[string][]urlEntry{
		"blog.urls": {{pattern: "posts/", viewName: "post_list"}},
	}
	// non-include: yields a single combined entry
	out := expandURLEntry(urlEntry{pattern: "x/", viewName: "v"}, "/api", byModule, map[string]bool{})
	if len(out) != 1 || out[0].pattern != "/api/x/" || out[0].viewName != "v" {
		t.Fatalf("non-include expand wrong: %+v", out)
	}
	// include: recurses into module with composed prefix
	out = expandURLEntry(urlEntry{pattern: "blog/", isInclude: true, includeModule: "blog.urls"}, "", byModule, map[string]bool{})
	if len(out) != 1 || out[0].pattern != "/blog/posts/" || out[0].viewName != "post_list" {
		t.Fatalf("include expand wrong: %+v", out)
	}
	// unresolved include: yields nothing
	out = expandURLEntry(urlEntry{pattern: "x/", isInclude: true, includeModule: "missing.urls"}, "", byModule, map[string]bool{})
	if len(out) != 0 {
		t.Fatalf("unresolved include should yield nothing, got %+v", out)
	}
}
