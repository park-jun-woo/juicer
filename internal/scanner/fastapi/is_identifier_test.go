//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestIsIdentifier 테스트
package fastapi

import "testing"

func TestIsIdentifier(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"API_STR", true},
		{"_private", true},
		{"abc123", true},
		{"", false},
		{"/api", false},
		{"123abc", false},
		{"a-b", false},
	}
	for _, tt := range tests {
		got := isIdentifier(tt.input)
		if got != tt.want {
			t.Errorf("isIdentifier(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
