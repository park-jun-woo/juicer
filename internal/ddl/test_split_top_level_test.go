//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestSplitTopLevel 테스트
package ddl

import (
	"testing"
)

func TestSplitTopLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   byte
		want  int
	}{
		{
			name:  "simple comma split",
			input: "a, b, c",
			sep:   ',',
			want:  3,
		},
		{
			name:  "comma inside parens ignored",
			input: "a, CHECK (x, y), b",
			sep:   ',',
			want:  3,
		},
		{
			name:  "no separator",
			input: "abc",
			sep:   ',',
			want:  1,
		},
		{
			name:  "empty input",
			input: "",
			sep:   ',',
			want:  0,
		},
		{
			name:  "whitespace only",
			input: "   ",
			sep:   ',',
			want:  0,
		},
		{
			name:  "nested parens",
			input: "a, b(c(d, e)), f",
			sep:   ',',
			want:  3,
		},
		{
			name:  "trailing whitespace only after last sep",
			input: "a, b,   ",
			sep:   ',',
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitTopLevel(tt.input, tt.sep)
			if len(got) != tt.want {
				t.Errorf("splitTopLevel() = %d parts, want %d; got: %v", len(got), tt.want, got)
			}
		})
	}
}
