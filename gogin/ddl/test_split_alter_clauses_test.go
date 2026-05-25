//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what TestSplitAlterClauses 테스트
package ddl

import (
	"testing"
)

func TestSplitAlterClauses(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single ADD COLUMN",
			input: "ADD COLUMN email TEXT",
			want:  1,
		},
		{
			name:  "ADD COLUMN and DROP COLUMN",
			input: "ADD COLUMN email TEXT, DROP COLUMN old_col",
			want:  2,
		},
		{
			name:  "single DROP COLUMN",
			input: "DROP COLUMN email",
			want:  1,
		},
		{
			name:  "ADD CONSTRAINT",
			input: "ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)",
			want:  1,
		},
		{
			name:  "DROP CONSTRAINT",
			input: "DROP CONSTRAINT fk_user",
			want:  1,
		},
		{
			name:  "ALTER COLUMN",
			input: "ALTER COLUMN id SET DEFAULT 0",
			want:  1,
		},
		{
			name:  "RENAME",
			input: "RENAME TO new_name",
			want:  1,
		},
		{
			name:  "continuation clause (column def with commas in default)",
			input: "ADD COLUMN data JSON DEFAULT '{\"a\": 1}'",
			want:  1,
		},
		{
			name:  "empty input",
			input: "",
			want:  0,
		},
		{
			name:  "multiple mixed clauses",
			input: "ADD COLUMN a TEXT, ADD COLUMN b INT, DROP COLUMN c",
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitAlterClauses(tt.input)
			if len(got) != tt.want {
				t.Errorf("splitAlterClauses() = %d clauses, want %d; got: %v", len(got), tt.want, got)
			}
		})
	}
}
