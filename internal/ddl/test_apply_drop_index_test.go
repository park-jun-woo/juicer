//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyDropIndex 테스트
package ddl

import (
	"testing"
)

func TestApplyDropIndex(t *testing.T) {
	t.Run("drop existing index", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Indexes: []string{"CREATE INDEX idx_users_name ON users (name)"},
			},
		}
		applyDropIndex(tables, "idx_users_name")
		if len(tables["users"].Indexes) != 0 {
			t.Errorf("expected 0 indexes, got %d", len(tables["users"].Indexes))
		}
	})

	t.Run("drop non-matching index", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Indexes: []string{"CREATE INDEX idx_users_email ON users (email)"},
			},
		}
		applyDropIndex(tables, "idx_users_name")
		if len(tables["users"].Indexes) != 1 {
			t.Errorf("expected 1 index, got %d", len(tables["users"].Indexes))
		}
	})

	t.Run("multiple tables", func(t *testing.T) {
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Indexes: []string{"CREATE INDEX idx_shared ON users (name)"},
			},
			"posts": {
				Name:    "posts",
				Indexes: []string{"CREATE INDEX idx_posts_title ON posts (title)"},
			},
		}
		applyDropIndex(tables, "idx_shared")
		if len(tables["users"].Indexes) != 0 {
			t.Error("expected idx_shared removed from users")
		}
		if len(tables["posts"].Indexes) != 1 {
			t.Error("expected idx_posts_title to remain")
		}
	})
}
