//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestStripInlineComments 테스트
package ddl

import (
	"testing"
)

func TestStripInlineComments(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no comments",
			input: "SELECT 1",
			want:  "SELECT 1",
		},
		{
			name:  "inline comment",
			input: "id BIGINT -- primary key",
			want:  "id BIGINT ",
		},
		{
			name:  "multiple lines with comments",
			input: "id BIGINT, -- pk\nname TEXT -- user name",
			want:  "id BIGINT, \nname TEXT ",
		},
		{
			name:  "no comment on some lines",
			input: "id BIGINT,\nname TEXT -- comment",
			want:  "id BIGINT,\nname TEXT ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripInlineComments(tt.input)
			if got != tt.want {
				t.Errorf("stripInlineComments() = %q, want %q", got, tt.want)
			}
		})
	}
}
