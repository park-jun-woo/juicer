//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestSplitStatements 테스트
package ddl

import (
	"testing"
)

func TestSplitStatements(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single statement",
			input: "SELECT 1;",
			want:  1,
		},
		{
			name:  "two statements",
			input: "SELECT 1; SELECT 2;",
			want:  2,
		},
		{
			name:  "empty input",
			input: "",
			want:  0,
		},
		{
			name:  "no trailing semicolon",
			input: "SELECT 1",
			want:  1,
		},
		{
			name:  "semicolon inside parentheses",
			input: "CREATE TABLE t (id INT; name TEXT);",
			// The semicolon inside parens should be kept, treated as single statement
			want: 1,
		},
		{
			name:  "multiple with whitespace",
			input: "  SELECT 1;  \n  SELECT 2;  \n",
			want:  2,
		},
		{
			name:  "trailing content without semicolon",
			input: "SELECT 1; SELECT 2",
			want:  2,
		},
		{
			name:  "only whitespace",
			input: "   \n\t  ",
			want:  0,
		},
		{
			name:  "nested parentheses with semicolons",
			input: "CREATE TABLE t (id INT, CHECK (x IN (1;2)));",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitStatements(tt.input)
			if len(got) != tt.want {
				t.Errorf("splitStatements() = %d statements, want %d; got: %v", len(got), tt.want, got)
			}
		})
	}
}
