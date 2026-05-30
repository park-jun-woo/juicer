//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what extractParamName 테스트
package fastify

import "testing"

func TestExtractParamName(t *testing.T) {
	tests := []struct{ in, want string }{
		{":id", "id"},
		{":id(\\d+)", "id"}, // regex constraint stripped
		{"users", ""},       // not a param
		{":", ""},           // empty name
	}
	for _, tt := range tests {
		if got := extractParamName(tt.in); got != tt.want {
			t.Errorf("extractParamName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
