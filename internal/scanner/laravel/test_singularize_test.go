//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what 영어 단수화 테스트
package laravel

import "testing"

func TestSingularize(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"users", "user"},
		{"posts", "post"},
		{"categories", "category"},
		{"boxes", "box"},
		{"classes", "class"},
		{"status", "status"},
	}
	for _, tt := range tests {
		got := singularize(tt.in)
		if got != tt.want {
			t.Errorf("singularize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
