//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyCreateTable_WithConstraintsCov 테스트
package ddl

import "testing"

func TestApplyCreateTable_WithConstraintsCov(t *testing.T) {
	tables := make(map[string]*Table)
	stmt := "CREATE TABLE orders (id INT, total DECIMAL, CONSTRAINT pk PRIMARY KEY (id))"
	applyCreateTable(tables, "orders", stmt)
	tbl := tables["orders"]
	if tbl == nil {
		t.Fatal("expected table")
	}
	if len(tbl.Constraints) < 1 {
		t.Fatal("expected at least one constraint")
	}
}
