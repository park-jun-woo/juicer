//ff:func feature=sql type=parse control=iteration dimension=1
//ff:what TestExtractTablesFromSQL 테스트
package sqls

import (
	"testing"
)

func TestExtractTablesFromSQL(t *testing.T) {
	tests := []struct {
		sql  string
		want int
	}{
		{"SELECT id FROM users", 1},
		{"INSERT INTO orders (id) VALUES (1)", 1},
		{"UPDATE users SET name = 'test' WHERE id = 1", 1},
		{"DELETE FROM users WHERE id = 1", 1},
		{"SELECT * FROM users JOIN orders ON users.id = orders.user_id", 2},
		{"SELECT 1", 0},
		{"SELECT * FROM lateral", 0}, // reserved word should be excluded
	}

	for _, tt := range tests {
		got := extractTablesFromSQL(tt.sql)
		if len(got) != tt.want {
			t.Errorf("extractTablesFromSQL(%q) = %d tables, want %d: %v", tt.sql, len(got), tt.want, got)
		}
	}
}
