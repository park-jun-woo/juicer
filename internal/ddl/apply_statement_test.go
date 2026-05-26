//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_CreateTable 테스트
package ddl

import "testing"

func TestApplyStatement_CreateTable(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "CREATE TABLE users (id INT PRIMARY KEY)")
	if tables["users"] == nil {
		t.Fatal("expected users table")
	}

	// CREATE INDEX
	applyStatement(tables, "CREATE INDEX idx_name ON users (name)")
	if len(tables["users"].Indexes) != 1 {
		t.Fatal("expected 1 index")
	}

	// DROP INDEX
	applyStatement(tables, "DROP INDEX idx_name")
	if len(tables["users"].Indexes) != 0 {
		t.Fatal("expected 0 indexes after drop")
	}

	// ALTER TABLE
	applyStatement(tables, "ALTER TABLE users ADD COLUMN email TEXT")
	if len(tables["users"].Columns) != 2 {
		t.Fatal("expected 2 columns after alter")
	}

	// DROP TABLE
	applyStatement(tables, "DROP TABLE users")
	if tables["users"] != nil {
		t.Fatal("expected users table to be dropped")
	}

	// unrecognized statement (no-op)
	applyStatement(tables, "INSERT INTO foo VALUES (1)")
}
