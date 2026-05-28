//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what 접두사와 경로를 결합한다
package django

import "testing"

func TestCombinePath(t *testing.T) {
	tests := []struct {
		prefix string
		path   string
		want   string
	}{
		{"/api", "users/", "/api/users/"},
		{"/api/", "users/", "/api/users/"},
		{"", "users/", "/users/"},
		{"/api", "", "/api"},
		{"", "", "/"},
	}

	for _, tt := range tests {
		got := combinePath(tt.prefix, tt.path)
		if got != tt.want {
			t.Errorf("combinePath(%q, %q) = %q, want %q", tt.prefix, tt.path, got, tt.want)
		}
	}
}
