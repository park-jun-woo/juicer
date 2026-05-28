//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what Hono :param → OpenAPI {param} 변환 테스트
package hono

import "testing"

func TestHonoPathToOpenAPI(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"/users/:id", "/users/{id}"},
		{"/users/:id/posts/:postId", "/users/{id}/posts/{postId}"},
		{"/users", "/users"},
		{"/:id", "/{id}"},
	}
	for _, tc := range tests {
		got := honoPathToOpenAPI(tc.input)
		if got != tc.expected {
			t.Errorf("honoPathToOpenAPI(%q) = %q, want %q", tc.input, got, tc.expected)
		}
	}
}
