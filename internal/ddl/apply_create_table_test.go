//ff:func feature=ddl type=test control=sequence
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

	// empty body
	applyCreateTable(tables, "empty", "CREATE TABLE empty")
	if tables["empty"] == nil {
		t.Fatal("expected empty table")
	}

	// constraint line
	tables2 := make(map[string]*Table)
	applyCreateTable(tables2, "orders", "CREATE TABLE orders (id INT, total INT, CONSTRAINT pk PRIMARY KEY (id))")
	if len(tables2["orders"].Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tables2["orders"].Constraints))
	}

	// empty line and extractColumnName returns ""
	tables3 := make(map[string]*Table)
	applyCreateTable(tables3, "t", "CREATE TABLE t (id INT, , )")
	if tables3["t"] == nil {
		t.Fatal("expected table t")
	}
}
