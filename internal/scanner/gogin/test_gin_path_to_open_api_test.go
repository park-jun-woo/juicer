//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestGinPathToOpenAPI 테스트
package gogin

import (
	"testing"
)

func TestGinPathToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"/users/:id", "/users/{id}"},
		{"/files/*path", "/files/{path}"},
		{"/users", "/users"},
		{"/users/:id/posts/:postId", "/users/{id}/posts/{postId}"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ginPathToOpenAPI(tt.input)
			if got != tt.want {
				t.Errorf("ginPathToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
