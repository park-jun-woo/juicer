//ff:func feature=ddl type=test control=sequence
//ff:what TestStripColumnReferences_Match 테스트
package ddl

import "testing"

func TestStripColumnReferences_Match(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "id", Raw: "id BIGINT"},
			{Name: "owner_id", Raw: "owner_id BIGINT NOT NULL REFERENCES owners(id) ON DELETE CASCADE"},
		},
	}
	stripColumnReferences(tbl, "owner_id")

	if tbl.Columns[1].Raw != "owner_id BIGINT NOT NULL" {
		t.Fatalf("expected inline references stripped, got %q", tbl.Columns[1].Raw)
	}

	if tbl.Columns[0].Raw != "id BIGINT" {
		t.Fatalf("unrelated column changed: %q", tbl.Columns[0].Raw)
	}
}
