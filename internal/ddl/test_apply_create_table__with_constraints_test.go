//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateTable_WithConstraints 테스트
package ddl

import "testing"

func TestApplyCreateTable_WithConstraints(t *testing.T) {
	tables := make(map[string]*Table)
	stmt := "CREATE TABLE users (id INT, name TEXT, FOREIGN KEY (id) REFERENCES other(id))"
	applyCreateTable(tables, "users", stmt)
	tbl := tables["users"]
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
}
