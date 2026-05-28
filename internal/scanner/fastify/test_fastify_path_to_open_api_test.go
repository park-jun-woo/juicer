//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what Fastify 경로 -> OpenAPI 경로 변환 테스트
package fastify

import "testing"

func TestFastifyPathToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/users/:id", "/users/{id}"},
		{"/users/:userId/posts/:postId", "/users/{userId}/posts/{postId}"},
		{"/users", "/users"},
		{"/:id", "/{id}"},
	}
	for _, tt := range tests {
		got := fastifyPathToOpenAPI(tt.input)
		if got != tt.want {
			t.Errorf("fastifyPathToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
