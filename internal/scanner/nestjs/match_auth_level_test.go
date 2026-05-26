//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestMatchAuthLevel 테스트
package nestjs

import "testing"

func TestMatchAuthLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"ApiAuth", "ApiAuth", "auth_required"},
		{"ApiPublic", "ApiPublic", "public"},
		{"Public", "Public", "public"},
		{"AuthOptional", "AuthOptional", "auth_optional"},
		{"plain Get", "Get", ""},
		{"plain Post", "Post", ""},
		{"HttpCode", "HttpCode", ""},
	}
	for _, tt := range tests {
		got := matchAuthLevel(tt.input)
		if got != tt.want {
			t.Errorf("matchAuthLevel(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
