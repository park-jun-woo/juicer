//ff:func feature=scan type=test control=sequence topic=django
//ff:what include를 따라 모듈을 재귀 전개하며 prefix를 합성하는지 검증한다
package django

import "testing"

func TestExpandURLModule(t *testing.T) {
	byModule := map[string][]urlEntry{
		"config.urls": {
			{pattern: "blog/", isInclude: true, includeModule: "blog.urls"},
		},
		"blog.urls": {
			{pattern: "posts/<int:pk>/", viewName: "post_detail"},
		},
	}
	entries := expandURLModule("config.urls", "", byModule, map[string]bool{})
	if len(entries) != 1 {
		t.Fatalf("expected 1 expanded entry, got %d", len(entries))
	}
	if entries[0].pattern != "/blog/posts/<int:pk>/" {
		t.Errorf("prefix composition wrong: got %q", entries[0].pattern)
	}
	if entries[0].viewName != "post_detail" {
		t.Errorf("viewName lost: got %q", entries[0].viewName)
	}
}
