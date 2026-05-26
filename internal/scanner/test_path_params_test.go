//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestPathParams 테스트
package scanner

import (
	"testing"
)

func TestPathParams(t *testing.T) {
	tests := []struct {
		path string
		want int
	}{
		{"/users/:id", 1},
		{"/users/:id/posts/:postId", 2},
		{"/users", 0},
	}
	for _, tt := range tests {
		got := PathParams(tt.path)
		if len(got) != tt.want {
			t.Errorf("PathParams(%q) = %d params, want %d", tt.path, len(got), tt.want)
		}
	}
}
