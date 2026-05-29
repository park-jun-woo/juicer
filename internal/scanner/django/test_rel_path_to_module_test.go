//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what 파일 상대경로를 dotted 모듈 경로로 변환하는지 검증한다
package django

import "testing"

func TestRelPathToModule(t *testing.T) {
	tests := []struct {
		rel  string
		want string
	}{
		{"blog/urls.py", "blog.urls"},
		{"urls.py", "urls"},
		{"pkg/__init__.py", "pkg"},
		{"a/b/c/urls.py", "a.b.c.urls"},
	}
	for _, tt := range tests {
		if got := relPathToModule(tt.rel); got != tt.want {
			t.Errorf("relPathToModule(%q) = %q, want %q", tt.rel, got, tt.want)
		}
	}
}
