//ff:func feature=sql type=parse control=sequence
//ff:what TestExtractTables 테스트
package sqls

import (
	"testing"
)

func TestExtractTables(t *testing.T) {
	fragments := []string{
		"SELECT id FROM users",
		"INSERT INTO orders (user_id) VALUES ($1)",
		"SELECT * FROM users JOIN orders ON users.id = orders.user_id",
	}

	tables := extractTables(fragments)
	if len(tables) < 2 {
		t.Errorf("expected at least 2 tables, got %d: %v", len(tables), tables)
	}
}
