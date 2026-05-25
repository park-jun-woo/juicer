//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestRefineCRUD 테스트
package sqls

import (
	"testing"
)

func TestRefineCRUD(t *testing.T) {
	tests := []struct {
		fragments []string
		want      string
	}{
		{[]string{"INSERT INTO users VALUES ($1)"}, "INSERT"},
		{[]string{"UPDATE users SET name = $1"}, "UPDATE"},
		{[]string{"DELETE FROM users WHERE id = $1"}, "DELETE"},
		{[]string{"SELECT * FROM users"}, "EXEC"},
		{nil, "EXEC"},
	}

	for _, tt := range tests {
		got := refineCRUD(tt.fragments)
		if got != tt.want {
			t.Errorf("refineCRUD(%v) = %q, want %q", tt.fragments, got, tt.want)
		}
	}
}
