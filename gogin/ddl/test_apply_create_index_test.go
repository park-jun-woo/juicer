//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateIndex 테스트
package ddl

import (
	"testing"
)

func TestApplyCreateIndex(t *testing.T) {
	t.Run("add index to existing table", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {Name: "users"},
		}
		applyCreateIndex(tables, "users", "CREATE INDEX idx_users_name ON users (name)")
		if len(tables["users"].Indexes) != 1 {
			t.Errorf("expected 1 index, got %d", len(tables["users"].Indexes))
		}
	})

	t.Run("add index to nonexistent table", func(t *testing.T) {
		tables := make(map[string]*Table)
		// Should not panic
		applyCreateIndex(tables, "nonexistent", "CREATE INDEX idx ON nonexistent (col)")
	})
}
