package ddl

import "testing"

func TestApplyCreateIndex_Existing(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users"},
	}
	applyCreateIndex(tables, "users", "CREATE INDEX idx_users_name ON users (name)")
	if len(tables["users"].Indexes) != 1 {
		t.Fatalf("expected 1 index, got %d", len(tables["users"].Indexes))
	}
}

func TestApplyCreateIndex_NoTable(t *testing.T) {
	tables := map[string]*Table{}
	applyCreateIndex(tables, "nonexistent", "CREATE INDEX idx ON nonexistent (x)")
}
