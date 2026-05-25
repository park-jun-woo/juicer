//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestExtractColumnName 테스트
package ddl

import (
	"testing"
)

func TestExtractColumnName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple column",
			input: "id BIGINT PRIMARY KEY",
			want:  "id",
		},
		{
			name:  "with leading whitespace",
			input: "  name TEXT NOT NULL",
			want:  "name",
		},
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "only whitespace",
			input: "   ",
			want:  "",
		},
		{
			name:  "uppercase becomes lowercase",
			input: "ID BIGINT",
			want:  "id",
		},
		{
			name:  "leading comment then column",
			input: "-- comment\nid BIGINT",
			want:  "id",
		},
		{
			name:  "only comment no newline",
			input: "-- comment only",
			want:  "",
		},
		{
			name:  "multiple leading comments",
			input: "-- comment1\n-- comment2\nname TEXT",
			want:  "name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractColumnName(tt.input)
			if got != tt.want {
				t.Errorf("extractColumnName() = %q, want %q", got, tt.want)
			}
		})
	}
}
