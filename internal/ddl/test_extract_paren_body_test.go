//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestExtractParenBody 테스트
package ddl

import (
	"testing"
)

func TestExtractParenBody(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple parens",
			input: "(id BIGINT, name TEXT)",
			want:  "id BIGINT, name TEXT",
		},
		{
			name:  "no parens",
			input: "CREATE TABLE t",
			want:  "",
		},
		{
			name:  "nested parens",
			input: "(id BIGINT, CHECK (x > 0))",
			want:  "id BIGINT, CHECK (x > 0)",
		},
		{
			name:  "prefix before parens",
			input: "CREATE TABLE t (id INT)",
			want:  "id INT",
		},
		{
			name:  "unclosed paren",
			input: "(id BIGINT, name TEXT",
			want:  "id BIGINT, name TEXT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractParenBody(tt.input)
			if got != tt.want {
				t.Errorf("extractParenBody() = %q, want %q", got, tt.want)
			}
		})
	}
}
