//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateTable_Basic 테스트
package ddl

import "testing"

func TestApplyCreateTable_Basic(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "users", "CREATE TABLE users (id INT PRIMARY KEY, name TEXT NOT NULL)")
	tbl := tables["users"]
	if tbl == nil {
		t.Fatal("expected table")
	}
	if len(tbl.Columns) != 2 {
		t.Fatalf("expected 2 columns, got %d", len(tbl.Columns))
	}
}
