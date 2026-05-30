//ff:func feature=scan type=test control=sequence topic=flask
//ff:what combinePath 테스트
package flask

import "testing"

func TestCombinePath(t *testing.T) {
	tests := []struct{ prefix, path, want string }{
		{"/api", "/users", "/api/users"},
		{"/api/", "users", "/api/users"}, // trailing slash trimmed + leading added
		{"/api", "", "/api"},             // empty path
		{"/api", "/", "/api"},            // root path
		{"", "", "/"},                    // empty prefix + empty path
		{"", "/x", "/x"},                 // empty prefix
	}
	for _, tt := range tests {
		if got := combinePath(tt.prefix, tt.path); got != tt.want {
			t.Errorf("combinePath(%q,%q) = %q, want %q", tt.prefix, tt.path, got, tt.want)
		}
	}
}
