//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestStripLeadingComments 테스트
package ddl

import (
	"testing"
)

func TestStripLeadingComments(t *testing.T) {
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
			name:  "leading comment",
			input: "-- comment\nSELECT 1",
			want:  "SELECT 1",
		},
		{
			name:  "multiple leading comments",
			input: "-- comment 1\n-- comment 2\nSELECT 1",
			want:  "SELECT 1",
		},
		{
			name:  "empty lines then comment then code",
			input: "\n\n-- comment\nSELECT 1",
			want:  "SELECT 1",
		},
		{
			name:  "only comments",
			input: "-- comment 1\n-- comment 2",
			want:  "",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			name:  "comment after code preserved",
			input: "SELECT 1 -- inline comment",
			want:  "SELECT 1 -- inline comment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripLeadingComments(tt.input)
			if got != tt.want {
				t.Errorf("stripLeadingComments() = %q, want %q", got, tt.want)
			}
		})
	}
}
