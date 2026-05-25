//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestCleanLine 테스트
package ddl

import (
	"testing"
)

func TestCleanLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no comment",
			input: "id BIGINT PRIMARY KEY",
			want:  "id BIGINT PRIMARY KEY",
		},
		{
			name:  "with inline comment",
			input: "id BIGINT -- primary key",
			want:  "id BIGINT",
		},
		{
			name:  "leading/trailing whitespace",
			input: "  name TEXT  ",
			want:  "name TEXT",
		},
		{
			name:  "comment with leading whitespace",
			input: "  id BIGINT -- pk  ",
			want:  "id BIGINT",
		},
		{
			name:  "no space before --",
			input: "id BIGINT--pk",
			want:  "id BIGINT--pk",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanLine(tt.input)
			if got != tt.want {
				t.Errorf("cleanLine() = %q, want %q", got, tt.want)
			}
		})
	}
}
