//ff:func feature=ddl type=test control=sequence
//ff:what TestStripColumnReferences_CaseInsensitiveName 테스트
package ddl

import "testing"

func TestStripColumnReferences_CaseInsensitiveName(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "owner_id", Raw: "owner_id BIGINT REFERENCES owners(id)"},
		},
	}

	stripColumnReferences(tbl, "OWNER_ID")
	if tbl.Columns[0].Raw != "owner_id BIGINT" {
		t.Fatalf("expected stripped, got %q", tbl.Columns[0].Raw)
	}
}
