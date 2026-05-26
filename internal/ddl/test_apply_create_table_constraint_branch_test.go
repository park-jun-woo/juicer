//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyCreateTable_ConstraintBranch 테스트
package ddl

import "testing"

func TestApplyCreateTable_ConstraintBranch(t *testing.T) {
	tables := make(map[string]*Table)
	applyCreateTable(tables, "orders", "CREATE TABLE orders (id INT, total DECIMAL, CONSTRAINT pk_orders PRIMARY KEY (id))")
	tbl := tables["orders"]
	if tbl == nil {
		t.Fatal("expected table")
	}
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
}
