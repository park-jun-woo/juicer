//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestLeadingToken 테스트
package scanner

import "testing"

func TestLeadingToken(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"login", "login"},
		{"findAll", "find"},
		{"listUsers", "list"},
		{"add_article", "add"},
		{"Login", "login"},
		{"", ""},
		{"x", "x"},
		{"createUserProfile", "create"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := leadingToken(tt.input)
			if got != tt.want {
				t.Errorf("leadingToken(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
